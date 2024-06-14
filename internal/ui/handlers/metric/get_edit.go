package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования метрики
func handleGetEdit() error {

	// Получить шаблон для формы
	// для редактирования метрики
	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/edit.html",
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
		"GET /metrics/edit/{service}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.PathValue("service")
			name := r.PathValue("name")

			// Получить метрику
			data, err := metric(
				service,
				name,
			)
			if err != nil {
				// Заполнить шаблон ошибки 404
				notfoundTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон формы
			// для редактирования метрики
			editTmpl.ExecuteTemplate(
				w,
				"edit",
				data,
			)
		},
	)

	return nil
}
