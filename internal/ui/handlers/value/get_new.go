package value

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования значения
func handleGetNew() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/new.html",
		),
	)

	http.HandleFunc(
		"GET /values/new",
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
