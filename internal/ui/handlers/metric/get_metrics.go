package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы метрик
func handleGetMetrics() error {

	// Получить шаблон
	// для страницы метрик
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

			// Получить все метрики
			data, err := metrics()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для страницы метрик
			tmpl.Execute(
				w,
				data,
			)
		},
	)

	return nil
}
