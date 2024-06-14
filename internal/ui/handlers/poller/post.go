package poller

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания опрашивающего плагина
func handlePost() error {

	pollerTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/poller/poller.html",
		),
	)

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

			service := r.FormValue("service")
			poller.Service = service

			metric := r.FormValue("metric")
			poller.Metric = metric

			err := insert(
				poller,
			)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			pollerTmpl.ExecuteTemplate(
				w,
				"poller",
				poller,
			)
		},
	)

	return nil
}
