package app

import (
	"github.com/khussa1n/Goods/app_sender/internal/config"
	"github.com/khussa1n/Goods/app_sender/internal/handler"
	"github.com/khussa1n/Goods/app_sender/internal/natspkg"
	"github.com/khussa1n/Goods/app_sender/internal/repository/pgrepo"
	"github.com/khussa1n/Goods/app_sender/internal/repository/redisrepo"
	"github.com/khussa1n/Goods/app_sender/internal/service"
	"github.com/khussa1n/Goods/app_sender/pkg/client/postgres"
	"github.com/khussa1n/Goods/app_sender/pkg/client/redis"
	"github.com/khussa1n/Goods/app_sender/pkg/httpserver"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {
	// Соеденение с базой
	pgConn, err := postgres.New(
		postgres.WithHost(cfg.DB.Postgres.Host),
		postgres.WithPort(cfg.DB.Postgres.Port),
		postgres.WithDBName(cfg.DB.Postgres.DBName),
		postgres.WithUsername(cfg.DB.Postgres.Username),
		postgres.WithPassword(cfg.DB.Postgres.Password),
	)
	if err != nil {
		log.Fatalf("connection to DB err: %s", err.Error())
		return err
	}
	log.Printf("connection to postgres success")

	migration := pgrepo.NewMigrate(cfg)
	err = migration.MigrateToVersion(cfg.DB.Postgres.MigrationVersion)
	if err != nil {
		log.Fatalf("from migration to postgres")
		return err
	}
	log.Printf("migration to postgres success")

	redisConn, err := redis.New(
		redis.WithHost(cfg.DB.Redis.Host),
		redis.WithPort(cfg.DB.Redis.Port),
		redis.WithDBName(cfg.DB.Redis.DBName),
	)
	if err != nil {
		log.Fatalf("connection to DB err: %s", err.Error())
		return err
	}
	log.Printf("connection to redis success")

	nc, err := natspkg.Connect(cfg.Nats.Host)
	if err != nil {
		log.Fatalf("connection to Nats err: %s", err.Error())
		return err
	}
	log.Printf("connection to nats success")

	postgresDB := pgrepo.New(pgConn.Pool)
	redisDB := redisrepo.New(redisConn.RedisClient)

	srvs := service.New(postgresDB, redisDB, cfg, nc)
	hndlr := handler.New(srvs)
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	// Запуск http сервера
	server.Start()
	log.Println("server started")

	// Создание канала для сигналов
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Ожидание сигналов от операционной системы и ошибки от http сервера
	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	// Принудительное остановление сервера
	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
