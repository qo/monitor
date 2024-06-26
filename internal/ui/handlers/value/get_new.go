package value

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования значения
func handleGetNew() error {

	// Получить шаблон
	// для формы
	// для создания нового значения
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/new.html",
		),
	)

	http.HandleFunc(
		"GET /values/new",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить все сервисы
			data, err := services()
			if err != nil {
				// TODO: использовать шаблон ошибки 500
			}

			// Заполнить шаблон
			// для формы
			// для создания нового значения
			tmpl.ExecuteTemplate(
				w,
				"new",
				data,
			)
		},
	)

	return nil
}
