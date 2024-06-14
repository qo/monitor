package endpoint

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// получения формы создания эндпоинта
func handleGetNew() error {

	type data struct {
		Messengers []db.Messenger
		Users      []db.User
	}

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/new.html",
		),
	)

	http.HandleFunc(
		"GET /endpoints/new",
		func(w http.ResponseWriter, r *http.Request) {
			var d data
			d.Messengers, err = messengers()
			if err != nil {
				// TODO: do something
			}
			d.Users, err = users()
			if err != nil {
				// TODO: do something
			}
			tmpl.ExecuteTemplate(
				w,
				"new",
				d,
			)
		},
	)

	return nil
}
