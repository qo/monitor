package poller

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
			"./internal/ui/templates/poller/metric.html",
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
		"GET /metrics-by-service-for-poller",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить параметры формы
			service := r.FormValue("service")

			// Получить метрики сервиса
			data, err := metrics(
				service,
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
