package value

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления значения
func handleDelete() error {

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /values/remove/{service}/{metric}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			service := r.PathValue("service")
			metric := r.PathValue("metric")
			name := r.PathValue("name")

			err := remove(service, metric, name)
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
