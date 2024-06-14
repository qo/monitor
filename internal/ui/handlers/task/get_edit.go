package task

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования задачи
func handleGetEdit() error {

	// Получить шаблон
	// для формы
	// для редактирования задачи
	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/task/edit.html",
		),
	)

	// Получить шаблон для ошибки 404
	notfoundTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/notfound.html",
		),
	)

	http.HandleFunc(
		"GET /tasks/edit/{service}/{metric}",
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
				// Заполнить шаблон для ошибки 404
				notfoundTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// для формы
			// для редактирования задачи
			editTmpl.ExecuteTemplate(
				w,
				"edit",
				data,
			)
		},
	)

	return nil
}
