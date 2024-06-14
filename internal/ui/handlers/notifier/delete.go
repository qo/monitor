package notifier

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления публикующего плагина
func handleDelete() error {

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /notifiers/remove/{messenger}",
		func(w http.ResponseWriter, r *http.Request) {

			messenger := r.PathValue("messenger")

			err := remove(
				messenger,
			)
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
