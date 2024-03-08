package chrepo

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/khussa1n/Goods/app_receiver/internal/entity"
	"io"
	"log"
	"os"
	"strings"
)

type Repo interface {
	InsertBatch(goodsBatch []entity.Goods) error
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

func (c *ClickhouseRepo) Migration(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open SQL file: %v", err)
	}
	defer file.Close()

	migrationContent, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %v", err)
	}

	err = c.DB.Exec(context.Background(), string(migrationContent))
	if err != nil {
		return err
	}
	return nil
}

func (c *ClickhouseRepo) InsertBatch(goodsBatch []entity.Goods) error {
	query := `
		INSERT INTO goods (Id, ProjectId, Name, Description, Priority, Removed, EventTime)
		VALUES `

	var placeholders string
	var values []interface{}

	for _, goods := range goodsBatch {
		placeholders += "(?, ?, ?, ?, ?, ?, ?),"

		values = append(values,
			goods.ID,
			goods.ProjectID,
			goods.Name,
			goods.Description,
			goods.Priority,
			goods.Removed,
			goods.EventTime,
		)
	}

	placeholders = strings.TrimSuffix(placeholders, ",")

	finalQuery := query + placeholders

	err := c.DB.Exec(context.Background(), finalQuery, values...)
	if err != nil {
		return err
	}

	log.Printf("Batch inserted successfully, size: %d", len(goodsBatch))

	return nil
}
