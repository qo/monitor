package trigger

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания триггера
func handlePost() error {

	triggerTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/trigger.html",
		),
	)

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

			service := r.FormValue("service")
			trigger.Service = service

			metric := r.FormValue("metric")
			trigger.Metric = metric

			value := r.FormValue("value")
			trigger.Value = value

			messenger := r.FormValue("messenger")
			trigger.Messenger = messenger

			user := r.FormValue("endpoint")
			trigger.User = user

			err := insert(trigger)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			triggerTmpl.ExecuteTemplate(
				w,
				"trigger",
				trigger,
			)
		},
	)

	return nil
}
