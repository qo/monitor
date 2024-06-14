package user

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания пользователя
func handlePost() error {

	userTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/user/user.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /user",
		func(w http.ResponseWriter, r *http.Request) {

			var user db.User

			username := r.FormValue("username")
			user.Username = username

			err := insert(user)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			userTmpl.ExecuteTemplate(
				w,
				"user",
				user,
			)
		},
	)

	return nil
}
