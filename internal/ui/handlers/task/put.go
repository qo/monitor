package task

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования задачи
func handlePut() error {

	// Получить шаблон
	// для формы с задачей
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/task/task.html",
		),
	)

	http.HandleFunc(
		"PUT /task/{service}/{metric}",
		func(w http.ResponseWriter, r *http.Request) {

			var task db.Task

			// Получения параметры формы
			service := r.FormValue("service")
			metric := r.FormValue("metric")
			value := r.FormValue("value")
			worker := r.FormValue("worker")

			// Заполнить данные о задаче
			task.Service = service
			task.Metric = metric
			task.Value = value
			task.Worker = worker

			// Обновить задачу
			err := update(task)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы с задачей
			tmpl.ExecuteTemplate(
				w,
				"task",
				task,
			)
		},
	)

	return nil
}
