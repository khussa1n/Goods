package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	HTTP ServerConfig `yaml:"http"`
	DB   DB           `yaml:"db"`
	Nats NatsConfig   `yaml:"nats"`
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
}

type DB struct {
	Postgres PostgresConfig `yaml:"postgres"`
	Redis    RedisConfig    `yaml:"redis"`
}

type PostgresConfig struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	DBName           string `yaml:"db_name"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	MigrationPath    string `yaml:"migration_path"`
	MigrationVersion uint   `yaml:"migration_version"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type NatsConfig struct {
	Host  string `yaml:"host"`
	Topic string `yaml:"topic"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
