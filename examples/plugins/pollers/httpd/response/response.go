package httpd_response

import (
	"errors"
	"fmt"
	"github.com/qo/monitor/examples/services/httpd/config"
	"io"
	"net/http"
)

// Данный плагин предназначен для того, чтобы
// возвращать значение метрики response
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
	return "response"
}

// Значения метрики response
const Error = "error"
const ValidResponse = "valid"
const InvalidResponse = "invalid"

// Метод, возвращающий текущее значение метрики
func (p poller) Value() (
	string,
	error,
) {

	errMsg := "can't get value of response metric of httpd service"

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
	// хост и порт и получить ответ
	resp, err := http.Get(
		fmt.Sprintf(
			"%s:%d",
			cfg.Host,
			cfg.Port,
		),
	)
	if err != nil {
		return Error,
			nil
	}

	// Прочитать тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Error,
			nil
	}

	// Если ответ совпал с желаемым,
	// вернуть значение ValidResponse
	if string(body) == cfg.Response {
		return ValidResponse,
			nil
	}

	// Вернуть значение InvalidResponse
	return InvalidResponse,
		nil
}

// Эксортировать переменную типа poller.
// Это можно назвать паттерном синглтон,
// так как тип poller не экспортирован,
// и как следствие переменные данного
// типа не могут быть созданы напрямую.
var Poller poller
