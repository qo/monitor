package service

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления сервиса
func handleDelete() error {

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /services/remove/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			name := r.PathValue("name")

			err := remove(name)
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
