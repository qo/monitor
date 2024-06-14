package metric

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания метрики
func handlePost() error {

	metricTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/metric.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /metric",
		func(w http.ResponseWriter, r *http.Request) {

			var metric db.Metric

			service := r.FormValue("service")
			metric.Service = service

			name := r.FormValue("name")
			metric.Name = name

			desc := r.FormValue("description")
			metric.Desc = desc

			err := insert(metric)
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
				metric,
			)
		},
	)

	return nil
}
