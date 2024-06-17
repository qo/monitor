package ui

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/qo/monitor/internal/config"
	"github.com/qo/monitor/internal/ui/handlers/daisy"
	"github.com/qo/monitor/internal/ui/handlers/endpoint"
	"github.com/qo/monitor/internal/ui/handlers/htmx"
	"github.com/qo/monitor/internal/ui/handlers/index"
	"github.com/qo/monitor/internal/ui/handlers/messenger"
	"github.com/qo/monitor/internal/ui/handlers/metric"
	"github.com/qo/monitor/internal/ui/handlers/notifier"
	"github.com/qo/monitor/internal/ui/handlers/poller"
	"github.com/qo/monitor/internal/ui/handlers/service"
	"github.com/qo/monitor/internal/ui/handlers/task"
	"github.com/qo/monitor/internal/ui/handlers/trigger"
	"github.com/qo/monitor/internal/ui/handlers/user"
	"github.com/qo/monitor/internal/ui/handlers/value"
)

// Зарегистрировать обработчики
// для пользовательского интерфейса
func Serve() error {

	errMsg := "can't serve ui: "

	var err error

	err = index.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = htmx.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = daisy.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = user.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = service.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = trigger.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = metric.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = value.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = messenger.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = endpoint.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = task.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = poller.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = notifier.Handle()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Загрузить конфигурацию
	config, err := config.Load()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	addr := fmt.Sprintf(
		"%s:%d",
		config.Host,
		config.Port,
	)

	// Запустить обработчики.
	// Данная функция завершится только
	// в случае, если появится исключение
	http.ListenAndServe(
		addr,
		nil,
	)

	return nil
}
