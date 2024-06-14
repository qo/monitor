package notifier

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы создания публикующего плагина
func handleGetNew() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/notifier/new.html",
		),
	)

	http.HandleFunc(
		"GET /notifiers/new",
		func(w http.ResponseWriter, r *http.Request) {

			messengers, err := messengers()
			if err != nil {
				// TODO: do something
			}

			tmpl.ExecuteTemplate(
				w,
				"new",
				messengers,
			)
		},
	)

	return nil
}
