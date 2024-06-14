package messenger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с мессенджером
func handleGetMessenger() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/messenger/messenger.html",
		),
	)

	http.HandleFunc(
		"GET /messenger/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			name := r.PathValue("name")
			data, err := messenger(
				name,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteTemplate(
				w,
				"messenger",
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
