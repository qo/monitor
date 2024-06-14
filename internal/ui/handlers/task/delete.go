package task

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления задачи
func handleDelete() error {

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /tasks/remove/{service}/{metric}",
		func(w http.ResponseWriter, r *http.Request) {

			service := r.PathValue("service")
			metric := r.PathValue("metric")

			err := remove(service, metric)
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
