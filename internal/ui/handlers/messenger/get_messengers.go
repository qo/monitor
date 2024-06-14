package messenger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы мессенджеров
func handleGetMessengers() error {

	// Шаблон страницы мессенджеров
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/messenger/messengers.html",
			"./internal/ui/templates/messenger/messenger.html",
			"./internal/ui/templates/messenger/new.html",
		),
	)

	http.HandleFunc(
		"GET /messengers",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить мессенджеры
			data, err := messengers()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон страницы мессенджеров
			tmpl.Execute(
				w,
				data,
			)
		},
	)

	return nil
}
