package config

import (
	"encoding/json"
	"os"
)

// Конфигурация HTTP-сервера
type config struct {
	// Хост для HTTP-сервера
	Host string `json:"host"`
	// Порт для HTTP-сервера
	Port int `json:"port"`
	// Ответ HTTP-сервера
	Response string `json:"response"`
}

// Загрузить конфигурацию
func Load() (config, error) {

	var c config

	file, err := os.Open("examples/services/httpd/httpd.conf")
	if err != nil {
		return c, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		return c, err
	}

	return c, nil
}
