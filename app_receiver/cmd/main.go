package main

import (
	"github.com/khussa1n/Goods/app_receiver/internal/app"
	"github.com/khussa1n/Goods/app_receiver/internal/config"
)

func main() {
	// Инициализация кофигурации
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	// Запуск
	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
