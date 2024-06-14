package user

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с пользователем
func handleGetUser() error {

	// Получить шаблон
	// для формы с пользователем
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/user/user.html",
		),
	)

	http.HandleFunc(
		"GET /user/{username}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получение значений query-параметров
			username := r.PathValue("username")

			// Получение пользователя
			data, err := user(
				username,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы с пользователем
			err = tmpl.ExecuteTemplate(
				w,
				"user",
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
