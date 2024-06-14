package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования метрики
func handleGetEdit() error {

	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/edit.html",
		),
	)

	notfoundTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/notfound.html",
		),
	)

	http.HandleFunc(
		"GET /metrics/edit/{service}/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			service := r.PathValue("service")
			name := r.PathValue("name")
			data, err := metric(
				service,
				name,
			)
			if err != nil {
				notfoundTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}
			editTmpl.ExecuteTemplate(
				w,
				"edit",
				data,
			)
		},
	)

	return nil
}
