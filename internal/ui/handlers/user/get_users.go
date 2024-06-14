package user

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с пользователями
func handleGetUsers() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/user/users.html",
			"./internal/ui/templates/user/user.html",
			"./internal/ui/templates/user/new.html",
		),
	)

	http.HandleFunc(
		"GET /users",
		func(w http.ResponseWriter, r *http.Request) {
			data, err := users()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			tmpl.Execute(
				w,
				data,
			)
		},
	)

	return nil
}
