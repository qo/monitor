package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с метрикой
func handleGetMetric() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/metric.html",
		),
	)

	http.HandleFunc(
		"GET /metric/{service}/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			service := r.PathValue("service")
			name := r.PathValue("name")
			data, err := metric(
				service,
				name,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteTemplate(
				w,
				"metric",
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
