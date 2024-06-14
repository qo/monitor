package user

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления пользователя
func handleDelete() error {

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /users/remove/{username}",
		func(w http.ResponseWriter, r *http.Request) {

			username := r.PathValue("username")

			err := remove(username)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}
		},
	)

	return nil
}
