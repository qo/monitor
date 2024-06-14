package metric

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования метрики
func handlePut() error {

	// Получить шаблон
	// для формы с метрикой
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/metric.html",
		),
	)

	http.HandleFunc(
		"PUT /metric/{service}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			var metric db.Metric

			// Получить значения query-параметров
			service := r.FormValue("service")
			name := r.FormValue("name")
			desc := r.FormValue("description")

			// Заполнить данные о метрике
			metric.Service = service
			metric.Name = name
			metric.Desc = desc

			// Обновить метрику
			err := update(
				metric,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы с метрикой
			tmpl.ExecuteTemplate(
				w,
				"metric",
				metric,
			)
		},
	)

	return nil
}
