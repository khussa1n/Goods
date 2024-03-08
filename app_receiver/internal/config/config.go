package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Nats       NatsConfig       `yaml:"nats"`
	Clickhouse ClickhouseConfig `yaml:"clickhouse"`
}

type NatsConfig struct {
	Host      string `yaml:"host"`
	Topic     string `yaml:"topic"`
	BatchSize int    `yaml:"batch_size"`
}

type ClickhouseConfig struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	DBName        string `yaml:"db_name"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	MigrationPath string `yaml:"migration_path"`
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
