package user

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с пользователями
func handleGetUsers() error {

	// Получить шаблон
	// для страницы с пользователями
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/user/users.html",
			"./internal/ui/templates/user/user.html",
			"./internal/ui/templates/user/new.html",
		),
	)

	http.HandleFunc(
		"GET /users",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить всех пользователей
			data, err := users()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для страницы с пользователями
			tmpl.Execute(
				w,
				data,
			)
		},
	)

	return nil
}
