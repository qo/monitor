package task

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования задачи
func handlePut() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/task/task.html",
		),
	)

	http.HandleFunc(
		"PUT /task/{service}/{metric}",
		func(w http.ResponseWriter, r *http.Request) {

			var task db.Task

			service := r.FormValue("service")
			task.Service = service

			metric := r.FormValue("metric")
			task.Metric = metric

			value := r.FormValue("value")
			task.Value = value

			worker := r.FormValue("worker")
			task.Worker = worker

			err := update(task)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tmpl.ExecuteTemplate(
				w,
				"task",
				task,
			)
		},
	)

	return nil
}
