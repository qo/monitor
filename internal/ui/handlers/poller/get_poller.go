package poller

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с опрашивающим плагином
func handleGetPoller() error {

	// Получить шаблон
	// для формы с опрашивающим плагином
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/poller/poller.html",
		),
	)

	http.HandleFunc(
		"GET /pollers/{service}/{metric}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.PathValue("service")
			metric := r.PathValue("metric")

			// Получить опрашивающий плагин
			data, err := poller(
				service,
				metric,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы с опрашивающим плагином
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
