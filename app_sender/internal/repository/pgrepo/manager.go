package pgrepo

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	projectTable = "projects"
	goodTable    = "goods"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Postgres {
	return &Postgres{
		Pool: pool,
	}
}
