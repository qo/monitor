package trigger

import (
	"net/http"
	"text/template"
)

func handleGetValuesByMetric() error {

	valueTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/value.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"GET /values-by-metric",
		func(w http.ResponseWriter, r *http.Request) {

			service := r.FormValue("service")
			metric := r.FormValue("metric")

			data, err := values(
				service,
				metric,
			)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			valueTmpl.ExecuteTemplate(
				w,
				"value",
				data,
			)
		},
	)

	return nil
}
