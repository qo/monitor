package trigger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с метриками
// указанного сервиса
func handleGetMetricsByService() error {

	// Получить шаблон
	// для формы с метрикой
	metricTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/metric.html",
		),
	)

	// Получить шаблон ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"GET /metrics-by-service-for-trigger",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.FormValue("service")

			// Получить метрики данного сервиса
			data, err := metrics(service)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// для формы с метрикой
			metricTmpl.ExecuteTemplate(
				w,
				"metric",
				data,
			)
		},
	)

	return nil
}
