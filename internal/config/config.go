package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// Конфигурация для программы
type config struct {

	// Хост и порт
	// для пользовательского интерфейса
	Host string
	Port int

	// Путь к БД
	DatabasePath string

	// Задержка между итерациями ядра
	// (в секундах)
	Delay time.Duration
}

// Промежуточная конфигурация для программы
// из соответствующего JSON-файла
type configJson struct {

	// Хост и порт
	// для пользовательского интерфейса
	Host string `json:"host"`
	Port int    `json:"port"`

	// Путь к БД
	DatabasePath string `json:"database_path"`

	// Задержка между итерациями ядра
	// (в секундах)
	DelaySeconds int `json:"delay_seconds"`
}

// Загрузить конфигурацию
func Load() (config, error) {

	errMsg := "can't load config: "

	// Временная конфигурация,
	// соответствующая JSON-файлу
	var cj configJson

	// Конфигурация
	var c config

	// Открыть файл конфигурации
	file, err := os.Open("monitor.conf")
	if err != nil {
		return c,
			errors.New(
				errMsg +
					"can't open config: " +
					err.Error(),
			)
	}
	defer file.Close()

	// Разобрать файл конфигурации в структуру
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cj)
	if err != nil {
		return c,
			errors.New(
				errMsg +
					"can't parse config: " +
					err.Error(),
			)
	}

	// Переменные для конфигурации
	host := cj.Host
	port := cj.Port
	databasePath := cj.DatabasePath
	delay, err := time.ParseDuration(
		fmt.Sprintf(
			"%ds",
			cj.DelaySeconds,
		),
	)
	if err != nil {
		return c,
			errors.New(
				errMsg +
					"can't parse config: " +
					err.Error(),
			)
	}

	// Заполнить конфигурацию
	c = config{
		host,
		port,
		databasePath,
		delay,
	}

	return c, nil
}
