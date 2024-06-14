package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для создания новой метрики
func handleGetNew() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/new.html",
		),
	)

	http.HandleFunc(
		"GET /metrics/new",
		func(w http.ResponseWriter, r *http.Request) {
			data, err := services()
			if err != nil {
				// TODO: do something
			}
			tmpl.ExecuteTemplate(
				w,
				"new",
				data,
			)
		},
	)

	return nil
}
