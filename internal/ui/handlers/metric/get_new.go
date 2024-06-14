package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для создания новой метрики
func handleGetNew() error {

	// Получить шаблон
	// для формы
	// для создания новой метрики
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/metric/new.html",
		),
	)

	http.HandleFunc(
		"GET /metrics/new",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить все сервисы
			data, err := services()
			if err != nil {
				// TODO: использовать шаблон ошибки 500
			}

			// Заполнить шаблон
			// для формы
			// для создания новой метрики
			tmpl.ExecuteTemplate(
				w,
				"new",
				data,
			)
		},
	)

	return nil
}
