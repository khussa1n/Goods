package chrepo

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/khussa1n/Goods/app_receiver/internal/entity"
)

type Repo interface {
	Insert(receivedData *entity.Goods) error
}

type ClickhouseRepo struct {
	DB driver.Conn
}

func New(db driver.Conn) *ClickhouseRepo {
	return &ClickhouseRepo{
		DB: db,
	}
}

func (c *ClickhouseRepo) Insert(goods *entity.Goods) error {
	query := `
		INSERT INTO goods (id, project_id, name, description, priority, removed, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	args := []interface{}{
		goods.ID,
		goods.ProjectID,
		goods.Name,
		goods.Description,
		goods.Priority,
		goods.Removed,
		goods.EventTime,
	}

	_, err := c.DB.Query(context.Background(), query, args...)
	if err != nil {
		return err
	}

	return nil
}
