package clickhousepkg

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Ch struct {
	host     string
	username string
	password string
	port     string
	dbName   string
	Conn     driver.Conn
}

func New(opts ...Option) (*Ch, error) {
	p := new(Ch)

	for _, opt := range opts {
		opt(p)
	}

	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{fmt.Sprintf("%s:%s", p.host, p.port)},
			Auth: clickhouse.Auth{
				Database: p.dbName,
				Username: p.username,
				Password: p.password,
			},
		})
	)

	if err != nil {
		return nil, err
	}

	if err = conn.Ping(ctx); err != nil {
		if ex, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n",
				ex.Code, ex.Message, ex.StackTrace)
		}
		return nil, err
	}

	p.Conn = conn

	return p, nil
}
