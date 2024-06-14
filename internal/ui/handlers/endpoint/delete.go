package endpoint

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления эндпоинта
func handleDelete() error {

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /endpoints/remove/{messenger}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			messenger := r.PathValue("messenger")
			name := r.PathValue("name")

			err := remove(messenger, name)
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
