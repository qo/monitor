package messenger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления мессенджера
func handleDelete() error {

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /messengers/remove/{name}",
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
