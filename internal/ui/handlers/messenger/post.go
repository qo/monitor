package messenger

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания мессенджера
func handlePost() error {

	// Шаблон формы с мессенджером
	messengerTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/messenger/messenger.html",
		),
	)

	// Шаблон ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /messenger",
		func(w http.ResponseWriter, r *http.Request) {

			var messenger db.Messenger

			// Получить параметры формы
			name := r.FormValue("name")
			desc := r.FormValue("description")

			// Заполнить данные о мессенджере
			messenger.Name = name
			messenger.Desc = desc

			// Добавить мессенджер
			err := insert(
				messenger,
			)
			if err != nil {
				// Заполнить шаблон
				// с ошибкой 500
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// формы с мессенджером
			messengerTmpl.ExecuteTemplate(
				w,
				"messenger",
				messenger,
			)
		},
	)

	return nil
}
