package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с метриками
func handleGetMetrics() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/metric/metrics.html",
			"./internal/ui/templates/metric/metric.html",
			"./internal/ui/templates/metric/new.html",
		),
	)

	http.HandleFunc(
		"GET /metrics",
		func(w http.ResponseWriter, r *http.Request) {
			data, err := metrics()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			tmpl.Execute(
				w,
				data,
			)
		},
	)

	return nil
}
