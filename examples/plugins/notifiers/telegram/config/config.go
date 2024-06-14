package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Конфигурация публикующего плагина
// для Telegram
type config struct {
	// Токен Telegram-бота
	Token string
	// Частота опроса изменений Telegram-бота
	Timeout time.Duration
}

// Промежуточная конфигурация публикующего плагина
// из соответствующего JSON-файла
type configJson struct {
	// Частота опроса изменений Telegram-бота
	// (в секундах)
	TimeoutSeconds int `json:"timeout_seconds"`
}

func Load() (config, error) {

	var cj configJson
	var c config

	file, err := os.Open(
		"./examples/plugins/notifiers/telegram/telegram.conf",
	)
	if err != nil {
		return c, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cj)
	if err != nil {
		return c, err
	}

	// Получить таймаут типа time.Duration
	timeout, err := time.ParseDuration(
		fmt.Sprintf(
			"%ds",
			cj.TimeoutSeconds,
		),
	)

	// Прочитать токен Telegram-бота
	// из переменных окружения
	token := os.Getenv(
		"MONITOR_TELEGRAM_BOT_TOKEN",
	)

	c = config{
		Token:   token,
		Timeout: timeout,
	}

	return c, nil
}
