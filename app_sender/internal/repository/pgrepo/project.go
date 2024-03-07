package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/khussa1n/Goods/app_sender/internal/custom_error"
	"github.com/khussa1n/Goods/app_sender/internal/entity"
	"log"
)

func (postgres *Postgres) CreateProject(ctx context.Context, p *entity.Projects) (*entity.Projects, error) {
	project := new(entity.Projects)

	query := fmt.Sprintf(`
		INSERT INTO %s (
			name, -- 1
			created_at -- 2
		)
		VALUES ($1, $2) RETURNING *;
	`, projectTable)

	err := pgxscan.Get(ctx, postgres.Pool, project, query,
		p.Name, p.CreatedAt)
	if err != nil {
		return nil, err
	}

	log.Printf("Created new good with ID %d\n", project.ID)

	return project, nil
}

func (postgres *Postgres) GetAllProjects(ctx context.Context, limit int64, offset int64) (int64, []entity.Projects, error) {
	var (
		projects []entity.Projects
		total    int64
	)

	query := fmt.Sprintf(`
			SELECT * FROM %s order by id desc limit $1 offset $2
		`, projectTable)

	err := pgxscan.Select(ctx, postgres.Pool, &projects, query, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	queryCount := fmt.Sprintf(`
		SELECT count(id) as total FROM %s
	`, projectTable)

	err = pgxscan.Get(ctx, postgres.Pool, &total, queryCount)
	if err != nil {
		return 0, nil, err
	}

	log.Printf("Retrieved %d projects from the database\n", len(projects))

	return total, projects, nil

}

func (postgres *Postgres) DeleteProjectByID(ctx context.Context, id int64) error {
	tx, err := postgres.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf(`
		DELETE FROM %s WHERE id = $1;
	`, projectTable)

	_, err = postgres.Pool.Exec(ctx, query, id)
	if err == pgx.ErrNoRows {
		return custom_error.ErrProjectNotFound
	} else if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	log.Printf("Deleted project with ID %d\n", id)

	return nil
}

func (postgres *Postgres) UpdateProjectByID(ctx context.Context, id int64, name string) (*entity.Projects, error) {
	project := new(entity.Projects)

	tx, err := postgres.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf(`
		UPDATE %s SET name = $2  WHERE id = $1 RETURNING *;
	`, projectTable)

	err = pgxscan.Get(ctx, postgres.Pool, project, query, id, name)
	if err == pgx.ErrNoRows {
		return nil, custom_error.ErrProjectNotFound
	} else if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("Updated project with ID %d\n", id)

	return project, nil
}
