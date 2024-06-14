package trigger

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

type data struct {
	Services   []db.Service
	Messengers []db.Messenger
}

// Зарегистрировать обработчик для
// получения формы для создания триггера
func handleGetNew() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/new.html",
		),
	)

	http.HandleFunc(
		"GET /triggers/new",
		func(w http.ResponseWriter, r *http.Request) {
			var d data
			d.Services, err = services()
			if err != nil {
				// TODO: do something
			}
			d.Messengers, err = messengers()
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
