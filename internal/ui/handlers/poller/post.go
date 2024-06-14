package poller

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания опрашивающего плагина
func handlePost() error {

	// Получить шаблон
	// для формы с опрашивающим плагином
	pollerTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/poller/poller.html",
		),
	)

	// Получить шаблон для ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /poller",
		func(w http.ResponseWriter, r *http.Request) {

			var poller db.Poller

			// Получить параметры формы
			service := r.FormValue("service")
			metric := r.FormValue("metric")

			// Заполнить данные об опрашивающем плагине
			poller.Service = service
			poller.Metric = metric

			// Добавить опрашивающий плагин
			err := insert(
				poller,
			)
			if err != nil {
				// Заполнить шаблон для ошибки 500
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// для формы с опрашивающим плагином
			pollerTmpl.ExecuteTemplate(
				w,
				"poller",
				poller,
			)
		},
	)

	return nil
}
