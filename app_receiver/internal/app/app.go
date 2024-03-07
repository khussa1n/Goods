package app

import (
	"context"
	"github.com/khussa1n/Goods/app_receiver/internal/config"
	"github.com/khussa1n/Goods/app_receiver/internal/natspkg"
	"github.com/khussa1n/Goods/app_receiver/internal/repository/chrepo"
	"github.com/khussa1n/Goods/app_receiver/pkg/clickhousepkg"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {
	db, err := clickhousepkg.New(
		clickhousepkg.WithHost(cfg.Clickhouse.Host),
		clickhousepkg.WithPort(cfg.Clickhouse.Port),
		clickhousepkg.WithDBName(cfg.Clickhouse.DBName),
		clickhousepkg.WithUsername(cfg.Clickhouse.Username),
		clickhousepkg.WithPassword(cfg.Clickhouse.Password),
	)
	if err != nil {
		log.Fatalf("connection to clickhouse err: %s", err.Error())
		return err
	}
	log.Printf("connection to clickhouse success")

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

	err = db.Conn.Exec(context.Background(), migrationQuery)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration completed successfully")

	nats := natspkg.New(cfg.Nats.Host, cfg.Nats.Topic)
	natsConn, err := nats.Connect()
	if err != nil {
		log.Fatalf("connection to Nats err: %s", err.Error())
		panic(err)
	}
	log.Printf("connection to nats success")

	chrepo := chrepo.New(db.Conn)

	natsHandler := &natspkg.NatsHandler{
		Chrepo: chrepo,
	}

	// Подписка на тему NATS и установка обработчика
	sub, err := natsConn.Subscribe(nats.Topic, natsHandler.HandleMessage)
	if err != nil {
		log.Fatal("Error subscribing to NATS:", err)
		return err
	}
	defer sub.Unsubscribe()

	log.Println("server started")

	// Создание канала для сигналов
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	}

	return nil
}
