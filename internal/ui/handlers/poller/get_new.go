package poller

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для создания подписывающего плагина
func handleGetNew() error {

	// Получить шаблон
	// для формы
	// для создания нового подписывающего плагина
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/poller/new.html",
		),
	)

	http.HandleFunc(
		"GET /pollers/new",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить все сервисы
			data, err := services()
			if err != nil {
				// TODO: использовать шаблон ошибки 500
			}

			// Заполнить шаблон
			// для формы
			// для создания нового подписывающего плагина
			tmpl.ExecuteTemplate(
				w,
				"new",
				data,
			)
		},
	)

	return nil
}
