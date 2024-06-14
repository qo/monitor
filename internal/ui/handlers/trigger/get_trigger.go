package trigger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с триггером
func handleGetTrigger() error {

	// Получить шаблон
	// для формы с триггером
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/trigger.html",
		),
	)

	http.HandleFunc(
		"GET /triggers/{service}/{metric}/{value}/{messenger}/{endpoint}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.PathValue("service")
			metric := r.PathValue("metric")
			value := r.PathValue("value")
			messenger := r.PathValue("messenger")
			endpoint := r.PathValue("endpoint")

			// Получить триггер
			data, err := trigger(
				service,
				metric,
				value,
				messenger,
				endpoint,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы с триггером
			err = tmpl.ExecuteTemplate(
				w,
				"trigger",
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
