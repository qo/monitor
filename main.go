package main

import (
	"github.com/qo/monitor/internal/core"
	"github.com/qo/monitor/internal/ui"
)

// Точка входа
func main() {
	// Запустить в фоне обработчики для
	// пользовательского интерфейса
	go ui.Serve()

	// Запустить обработчик для ядра
	core.Run()
}
