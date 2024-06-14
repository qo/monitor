package trigger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления триггера
func handleDelete() error {

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /triggers/remove/{service}/{metric}/{value}/{messenger}/{endpoint}",
		func(w http.ResponseWriter, r *http.Request) {

			service := r.PathValue("service")
			metric := r.PathValue("metric")
			value := r.PathValue("value")
			messenger := r.PathValue("messenger")
			endpoint := r.PathValue("endpoint")

			err := remove(service, metric, value, messenger, endpoint)
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
