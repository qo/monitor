package poller

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с опрашивающим плагином
func handleGetPoller() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/poller/poller.html",
		),
	)

	http.HandleFunc(
		"GET /pollers/{service}/{metric}",
		func(w http.ResponseWriter, r *http.Request) {

			service := r.PathValue("service")
			metric := r.PathValue("metric")

			data, err := poller(
				service,
				metric,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(
				w,
				"poller",
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
