package task

import (
	"log"
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с задачами
func handleGetTasks() error {

	// Получить шаблон
	// страницы с задачами
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

			// Получить все задачи
			data, err := tasks()
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// страницы с задачами
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
