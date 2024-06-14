package user

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для создания пользователя
func handleGetNew() error {

	// Получить шаблон
	// для формы
	// для создания пользователя
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/user/new.html",
		),
	)

	http.HandleFunc(
		"GET /users/new",
		func(w http.ResponseWriter, r *http.Request) {
			// Заполнить шаблон
			// для формы
			// для создания пользователя
			tmpl.ExecuteTemplate(
				w,
				"new",
				nil,
			)
		},
	)

	return nil
}
