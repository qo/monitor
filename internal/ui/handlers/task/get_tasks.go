package task

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с задачами
func handleGetTasks() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/task/tasks.html",
			"./internal/ui/templates/task/task.html",
		),
	)

	http.HandleFunc(
		"GET /tasks",
		func(w http.ResponseWriter, r *http.Request) {
			data, err := tasks()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(
				w,
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
