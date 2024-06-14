package user

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с пользователем
func handleGetUser() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/user/user.html",
		),
	)

	http.HandleFunc(
		"GET /user/{username}",
		func(w http.ResponseWriter, r *http.Request) {
			username := r.PathValue("username")
			data, err := user(
				username,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteTemplate(
				w,
				"user",
				data,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)

	return nil
}
