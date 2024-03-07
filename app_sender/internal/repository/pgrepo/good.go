package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/khussa1n/Goods/app_sender/internal/custom_error"
	"github.com/khussa1n/Goods/app_sender/internal/entity"
	"github.com/khussa1n/Goods/app_sender/internal/entity/api"
	"log"
)

func (postgres *Postgres) CreateGood(ctx context.Context, g *entity.Goods) (*entity.Goods, error) {
	goods := new(entity.Goods)

	maxPriorityQuery := fmt.Sprintf(`
        SELECT COALESCE(MAX(priority), 0) FROM %s;
    `, goodTable)

	var maxPriority int64
	err := postgres.Pool.QueryRow(ctx, maxPriorityQuery).Scan(&maxPriority)
	if err != nil {
		return nil, err
	}

	newPriority := maxPriority + 1

	insertQuery := fmt.Sprintf(`
        INSERT INTO %s (
            name,
            project_id,
            description,
            priority,
            removed,
            created_at
        )
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;
    `, goodTable)

	err = pgxscan.Get(ctx, postgres.Pool, goods, insertQuery,
		g.Name, g.ProjectID, g.Description, newPriority, g.Removed, g.CreatedAt)
	if err != nil {
		return nil, err
	}

	return goods, nil
}

func (postgres *Postgres) GetAllGoods(ctx context.Context, limit int64, offset int64) (total int64, removed int64, goods []entity.Goods, err error) {
	query := fmt.Sprintf(`
        SELECT * FROM %s WHERE removed = false ORDER BY id DESC LIMIT $1 OFFSET $2
    `, goodTable)

	err = pgxscan.Select(ctx, postgres.Pool, &goods, query, limit, offset)
	if err != nil {
		return 0, 0, nil, err
	}

	queryCount := fmt.Sprintf(`
        SELECT count(id) as total FROM %s
    `, goodTable)

	err = pgxscan.Get(ctx, postgres.Pool, &total, queryCount)
	if err != nil {
		return 0, 0, nil, err
	}

	queryRemovedCount := fmt.Sprintf(`
        SELECT count(id) as removed FROM %s WHERE removed = true
    `, goodTable)

	err = pgxscan.Get(ctx, postgres.Pool, &removed, queryRemovedCount)
	if err != nil {
		return 0, 0, nil, err
	}

	return total, removed, goods, nil
}

func (postgres *Postgres) DeleteGoodByID(ctx context.Context, id int64) (*entity.Goods, error) {
	tx, err := postgres.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("DeleteGoodByID method, Rollback: %s \n", err.Error())
		}
	}(tx, ctx)

	good := new(entity.Goods)

	query := fmt.Sprintf(`
		UPDATE %s SET removed = true WHERE id = $1 AND removed = false RETURNING *;
	`, goodTable)

	_, err = postgres.Pool.Exec(ctx, query, id)
	if err == pgx.ErrNoRows {
		return nil, custom_error.ErrGoodNotFound
	} else if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return good, nil
}

func (postgres *Postgres) UpdateGoodByID(ctx context.Context, id int64, g *entity.Goods) (*entity.Goods, error) {
	goods := new(entity.Goods)

	tx, err := postgres.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("DeleteGoodByID method, Rollback: %s \n", err.Error())
		}
	}(tx, ctx)

	query := fmt.Sprintf(`
		UPDATE %s SET name = $2, description = $3 WHERE id = $1 RETURNING *;
	`, goodTable)

	err = pgxscan.Get(ctx, postgres.Pool, goods, query, id, g.Name, g.Description)
	if err == pgx.ErrNoRows {
		return nil, custom_error.ErrGoodNotFound
	} else if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return goods, nil
}

func (postgres *Postgres) Reprioritiize(ctx context.Context, id int64, priority int64) ([]api.Priorities, error) {
	tx, err := postgres.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("Reprioritize method, Rollback: %s \n", err.Error())
		}
	}(tx, ctx)

	updateQuery := fmt.Sprintf(`
		UPDATE %s SET priority = $2 WHERE id >= $1 RETURNING id, priority;
	`, goodTable)

	_, err = postgres.Pool.Exec(ctx, updateQuery, id, priority)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, custom_error.ErrGoodNotFound
		}
		return nil, err
	}

	rows, err := postgres.Pool.Query(ctx, updateQuery, id, priority)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, custom_error.ErrGoodNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var updatedPriorities []api.Priorities
	for rows.Next() {
		var priority api.Priorities
		if err := rows.Scan(&priority.Id, &priority.Priotiry); err != nil {
			return nil, err
		}
		updatedPriorities = append(updatedPriorities, priority)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return updatedPriorities, nil
}
