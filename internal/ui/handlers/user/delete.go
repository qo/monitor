package user

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления пользователя
func handleDelete() error {

	// Получить шаблон ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /users/remove/{username}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			username := r.PathValue("username")

			// Удалить пользователя
			err := remove(
				username,
			)
			if err != nil {
				// Заполнить шаблон ошибки 500
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
