package trigger

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания триггера
func handlePost() error {

	// Получить шаблон
	// для формы с триггером
	triggerTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/trigger.html",
		),
	)

	// Получить шаблон ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /trigger",
		func(w http.ResponseWriter, r *http.Request) {

			var trigger db.Trigger

			// Получить значения query-параметров
			service := r.FormValue("service")
			metric := r.FormValue("metric")
			value := r.FormValue("value")
			messenger := r.FormValue("messenger")
			user := r.FormValue("user")

			// Заполнить данные триггера
			trigger.Service = service
			trigger.Metric = metric
			trigger.Value = value
			trigger.Messenger = messenger
			trigger.User = user

			// Добавить триггер
			err := insert(
				trigger,
			)
			if err != nil {
				// Заполнить шаблон ошибки 500
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// для формы с триггером
			triggerTmpl.ExecuteTemplate(
				w,
				"trigger",
				trigger,
			)
		},
	)

	return nil
}
