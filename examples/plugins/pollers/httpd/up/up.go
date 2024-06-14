package httpd_up

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/qo/monitor/examples/services/httpd/config"
)

// Данный плагин предназначен для того, чтобы
// возвращать значение метрики up
// сервиса httpd

// Тип, который должен реализовать интерфейс
// plugins.Poller
type poller struct{}

// Метод, возвращающий имя сервиса,
// к которому относится плагин.
// Данный метод предназначен для того, чтобы
// связать данный код плагина с сущностью
// db.Poller.
func (p poller) Service() string {
	return "httpd"
}

// Аналогично методу Service()
func (p poller) Metric() string {
	return "up"
}

// Значения метрики up
const UpValue = "up"
const DownValue = "down"

// Метод, возвращающий текущее значение метрики
func (p poller) Value() (
	string,
	error,
) {

	errMsg := "can't get value of up metric of httpd service: "

	// Загрузить конфигурацию HTTP-сервера
	cfg, err := config.Load()
	if err != nil {
		return "",
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	// Послать GET-запрос на указанные
	// хост и порт
	_, err = http.Get(
		fmt.Sprintf(
			"%s:%d",
			cfg.Host,
			cfg.Port,
		),
	)
	if err != nil {
		return DownValue,
			nil
	}

	return UpValue,
		nil
}

// Эксортировать переменную типа poller.
// Это можно назвать паттерном синглтон,
// так как тип poller не экспортирован,
// и как следствие переменные данного
// типа не могут быть созданы напрямую.
var Poller poller
