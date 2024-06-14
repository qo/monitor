package endpoint

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// получения формы создания эндпоинта
func handleGetNew() error {

	// Данные для заполнения
	// шаблона для формы для создания нового эндпоинта
	type data struct {
		Messengers []db.Messenger
		Users      []db.User
	}

	// Шаблон для формы для создания нового эндпоинта
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/new.html",
		),
	)

	http.HandleFunc(
		"GET /endpoints/new",
		func(w http.ResponseWriter, r *http.Request) {

			var d data

			// Получить все мессенджеры
			d.Messengers, err = messengers()
			if err != nil {
				// TODO: использовать шаблон с ошибкой 500
			}

			// Получить всех пользователей
			d.Users, err = users()
			if err != nil {
				// TODO: использовать шаблон с ошибкой 500
			}

			// Заполнить шаблон данными
			tmpl.ExecuteTemplate(
				w,
				"new",
				d,
			)
		},
	)

	return nil
}
