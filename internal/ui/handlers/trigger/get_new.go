package trigger

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Данные для формы
type data struct {
	Services   []db.Service
	Messengers []db.Messenger
}

// Зарегистрировать обработчик для
// получения формы для создания триггера
func handleGetNew() error {

	// Получить шаблон
	// для формы
	// для создания триггера
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/new.html",
		),
	)

	http.HandleFunc(
		"GET /triggers/new",
		func(w http.ResponseWriter, r *http.Request) {

			var d data

			// Получить все сервисы
			d.Services, err = services()
			if err != nil {
				// TODO: использовать шаблон ошибки 500
			}

			// Получить все мессенджеры
			d.Messengers, err = messengers()
			if err != nil {
				// TODO: использовать шаблон ошибки 500
			}

			// Заполнить шаблон
			// для формы
			// для создания триггера
			tmpl.ExecuteTemplate(
				w,
				"new",
				d,
			)
		},
	)

	return nil
}
