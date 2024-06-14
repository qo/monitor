package metric

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания метрики
func handlePost() error {

	// Получить шаблон
	// для формы с метрикой
	metricTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/metric.html",
		),
	)

	// Получить шаблон
	// для ошибки 500
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

			// Получить параметры формы
			service := r.FormValue("service")
			name := r.FormValue("name")
			desc := r.FormValue("description")

			// Заполнить данные метрики
			metric.Service = service
			metric.Name = name
			metric.Desc = desc

			// Добавить метрику
			err := insert(
				metric,
			)
			if err != nil {
				// Заполнить шаблон ошибки 500
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
				metric,
			)
		},
	)

	return nil
}
