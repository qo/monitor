package messenger

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания мессенджера
func handlePost() error {

	messengerTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/messenger/messenger.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /messenger",
		func(w http.ResponseWriter, r *http.Request) {

			var messenger db.Messenger

			name := r.FormValue("name")
			messenger.Name = name

			desc := r.FormValue("description")
			messenger.Desc = desc

			err := insert(messenger)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			messengerTmpl.ExecuteTemplate(
				w,
				"messenger",
				messenger,
			)
		},
	)

	return nil
}
