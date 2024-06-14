package messenger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с мессенджерами
func handleGetMessengers() error {

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
			data, err := messengers()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			tmpl.Execute(
				w,
				data,
			)
		},
	)

	return nil
}
