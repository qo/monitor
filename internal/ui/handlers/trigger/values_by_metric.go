package trigger

import (
	"net/http"
	"text/template"
)

func handleGetValuesByMetric() error {

	// Получить шаблон
	// для формы со значением
	valueTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/value.html",
		),
	)

	// Получить шаблон для ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"GET /values-by-metric",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.FormValue("service")
			metric := r.FormValue("metric")

			// Получить значения данной метрики
			data, err := values(
				service,
				metric,
			)
			if err != nil {
				// Заполнить шаблон для ошибки 500
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// для формы со значением
			valueTmpl.ExecuteTemplate(
				w,
				"value",
				data,
			)
		},
	)

	return nil
}
