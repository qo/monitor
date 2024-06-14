package messenger

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования мессенджера
func handlePut() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/messenger/messenger.html",
		),
	)

	http.HandleFunc(
		"PUT /messenger/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			var messenger db.Messenger

			name := r.FormValue("name")
			messenger.Name = name

			desc := r.FormValue("description")
			messenger.Desc = desc

			err := update(messenger)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tmpl.ExecuteTemplate(
				w,
				"messenger",
				messenger,
			)
		},
	)

	return nil
}
