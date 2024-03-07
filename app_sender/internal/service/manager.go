package service

import (
	"github.com/khussa1n/Goods/app_sender/internal/config"
	"github.com/khussa1n/Goods/app_sender/internal/repository/pgrepo"
	"github.com/khussa1n/Goods/app_sender/internal/repository/redisrepo"
	"github.com/nats-io/nats.go"
)

type Manager struct {
	PostgresRepository pgrepo.PostgresRepository
	RedisRepository    redisrepo.RedisRepository
	Config             *config.Config
	NatsConn           *nats.Conn
}

func New(rostgresRepository pgrepo.PostgresRepository, redisRepository redisrepo.RedisRepository,
	config *config.Config, natsConn *nats.Conn) *Manager {
	return &Manager{
		PostgresRepository: rostgresRepository,
		RedisRepository:    redisRepository,
		Config:             config,
		NatsConn:           natsConn,
	}
}
