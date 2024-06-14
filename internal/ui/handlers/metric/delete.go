package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления метрики
func handleDelete() error {

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /metrics/remove/{service}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			service := r.PathValue("service")
			name := r.PathValue("name")

			err := remove(service, name)
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
