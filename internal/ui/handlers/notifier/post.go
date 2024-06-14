package notifier

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания публикующего плагина
func handlePost() error {

	// Получить шаблон
	// для формы с публикующим плагином
	notifierTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/notifier/notifier.html",
		),
	)

	// Получить шаблон
	// для ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /notifier",
		func(w http.ResponseWriter, r *http.Request) {

			var notifier db.Notifier

			// Получить параметры формы
			messenger := r.FormValue("messenger")

			// Заполнить данные публикующего плагина
			notifier.Messenger = messenger

			// Добавить публикующий плагин
			err := insert(
				notifier,
			)
			if err != nil {
				// Заполнить шаблон
				// для ошибки 500
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// для формы с публикующим плагином
			notifierTmpl.ExecuteTemplate(
				w,
				"notifier",
				notifier,
			)
		},
	)

	return nil
}
