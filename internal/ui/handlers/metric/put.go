package metric

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования метрики
func handlePut() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/metric.html",
		),
	)

	http.HandleFunc(
		"PUT /metric/{service}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			var metric db.Metric

			service := r.FormValue("service")
			metric.Service = service

			name := r.FormValue("name")
			metric.Name = name

			desc := r.FormValue("description")
			metric.Desc = desc

			err := update(metric)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tmpl.ExecuteTemplate(
				w,
				"metric",
				metric,
			)
		},
	)

	return nil
}
