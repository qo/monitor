package messenger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для создания мессенджера
func handleGetNew() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/messenger/new.html",
		),
	)

	http.HandleFunc(
		"GET /messengers/new",
		func(w http.ResponseWriter, r *http.Request) {
			if err != nil {
				// TODO: do something
			}
			tmpl.ExecuteTemplate(
				w,
				"new",
				nil,
			)
		},
	)

	return nil
}
