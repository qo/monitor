package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с метрикой
func handleGetMetric() error {

	// Получить шаблон
	// для формы с метрикой
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/metric.html",
		),
	)

	http.HandleFunc(
		"GET /metric/{service}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.PathValue("service")
			name := r.PathValue("name")

			// Получить метрику
			data, err := metric(
				service,
				name,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы с метрикой
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
