package notifier

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания публикующего плагина
func handlePost() error {

	notifierTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/notifier/notifier.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /notifier",
		func(w http.ResponseWriter, r *http.Request) {

			var notifier db.Notifier

			messenger := r.FormValue("messenger")
			notifier.Messenger = messenger

			err := insert(
				notifier,
			)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			notifierTmpl.ExecuteTemplate(
				w,
				"notifier",
				notifier,
			)
		},
	)

	return nil
}
