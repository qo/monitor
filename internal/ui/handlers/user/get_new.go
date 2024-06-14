package user

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для создания пользователя
func handleGetNew() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/user/new.html",
		),
	)

	http.HandleFunc(
		"GET /users/new",
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
