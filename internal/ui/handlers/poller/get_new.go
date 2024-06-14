package poller

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для создания опрашивающего плагина
func handleGetNew() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/poller/new.html",
		),
	)

	http.HandleFunc(
		"GET /pollers/new",
		func(w http.ResponseWriter, r *http.Request) {
			data, err := services()
			if err != nil {
				// TODO: do something
			}
			tmpl.ExecuteTemplate(
				w,
				"new",
				data,
			)
		},
	)

	return nil
}
