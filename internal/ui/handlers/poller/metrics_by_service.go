package poller

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с метриками
// указанного сервиса
func handleGetMetricsByService() error {

	metricTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/poller/metric.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"GET /metrics-by-service-for-poller",
		func(w http.ResponseWriter, r *http.Request) {

			service := r.FormValue("service")

			data, err := metrics(service)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			metricTmpl.ExecuteTemplate(
				w,
				"metric",
				data,
			)
		},
	)

	return nil
}
