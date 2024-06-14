package task

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с задачей
func handleGetTask() error {

	// Получить шаблон
	// для формы с задачей
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/task/task.html",
		),
	)

	http.HandleFunc(
		"GET /tasks/{service}/{metric}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.PathValue("service")
			metric := r.PathValue("metric")

			// Получить задачу
			data, err := task(
				service,
				metric,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы с задачей
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
