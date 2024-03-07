package chrepo

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/khussa1n/Goods/app_receiver/internal/entity"
	"log"
)

type Repo interface {
	Insert(receivedData *entity.Goods) error
	Migration() error
}

type ClickhouseRepo struct {
	DB driver.Conn
}

func New(db driver.Conn) *ClickhouseRepo {
	return &ClickhouseRepo{
		DB: db,
	}
}

func (c *ClickhouseRepo) Migration() error {
	migrationQuery := `
		CREATE TABLE IF NOT EXISTS goods (
			Id Int64,
			ProjectId Int64,
			Name String,
			Description String,
			Priority Int64,
			Removed UInt8,
			EventTime DateTime
		) ENGINE = MergeTree()
		ORDER BY (Id);
			`

	err := c.DB.Exec(context.Background(), migrationQuery)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClickhouseRepo) Insert(goods *entity.Goods) error {
	query := `
		INSERT INTO goods (Id, ProjectId, Name, Description, Priority, Removed, EventTime)
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

	err := c.DB.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}

	log.Println("Inserted successfully data: ", goods)

	return nil
}
