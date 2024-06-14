package user

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания пользователя
func handlePost() error {

	// Получить шаблон
	// для формы с пользователем
	userTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/user/user.html",
		),
	)

	// Получить шаблон для ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /user",
		func(w http.ResponseWriter, r *http.Request) {

			var user db.User

			// Получить параметры формы
			username := r.FormValue("username")

			// Заполнить данные о пользователе
			user.Username = username

			// Добавить пользователя
			err := insert(
				user,
			)
			if err != nil {
				// Заполнить шаблон ошибки 500
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// для формы с пользователем
			userTmpl.ExecuteTemplate(
				w,
				"user",
				user,
			)
		},
	)

	return nil
}
