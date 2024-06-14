package task

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с задачей
func handleGetTask() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/task/task.html",
		),
	)

	http.HandleFunc(
		"GET /tasks/{service}/{metric}",
		func(w http.ResponseWriter, r *http.Request) {
			service := r.PathValue("service")
			metric := r.PathValue("metric")
			data, err := task(
				service,
				metric,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteTemplate(
				w,
				"task",
				data,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)

	return nil
}
