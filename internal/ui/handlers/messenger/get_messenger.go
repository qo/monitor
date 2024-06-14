package messenger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с мессенджером
func handleGetMessenger() error {

	// Шаблон для
	// формы с мессенджером
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/messenger/messenger.html",
		),
	)

	http.HandleFunc(
		"GET /messenger/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			name := r.PathValue("name")

			// Получить мессенджер
			data, err := messenger(
				name,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон для
			// формы с мессенджером
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
