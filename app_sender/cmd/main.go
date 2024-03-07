package main

import (
	"github.com/khussa1n/Goods/app_sender/internal/app"
	"github.com/khussa1n/Goods/app_sender/internal/config"
)

// @title           Goods
// @version         0.0.1
// @description     API for Goods application.

// @contact.name   Khussain
// @contact.email  khussain.qudaibergenov@gmail.com

// @host      localhost:8081
// @BasePath  /
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
