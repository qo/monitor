package notifier

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы создания публикующего плагина
func handleGetNew() error {

	// Получить шаблон
	// для формы
	// для создания нового публикующего плагина
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/notifier/new.html",
		),
	)

	http.HandleFunc(
		"GET /notifiers/new",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить все мессенджеры
			messengers, err := messengers()
			if err != nil {
				// TODO: использовать шаблон ошибки 500
			}

			// Заполнить шаблон
			// для формы
			// для создания нового публикующего плагина
			tmpl.ExecuteTemplate(
				w,
				"new",
				messengers,
			)
		},
	)

	return nil
}
