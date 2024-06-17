package main

import (
	"errors"
	"log"
	"time"

	"github.com/qo/monitor/internal/config"
	"github.com/qo/monitor/internal/core"
	"github.com/qo/monitor/internal/ui"
)

// Точка входа
func main() {

	// Запустить в фоне обработчики для
	// пользовательского интерфейса
	go ui.Serve()

	// Загрузить конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(
			errors.New(
				"loading config failed: " +
					err.Error(),
			),
		)
	}

	// Запустить обработчик для ядра
	for {
		err := core.Run()
		log.Println(
			errors.New(
				"running core failed: " +
					err.Error(),
			),
		)
		time.Sleep(
			cfg.Delay,
		)
	}
}
