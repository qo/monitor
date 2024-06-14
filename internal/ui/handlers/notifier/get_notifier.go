package notifier

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с публикующим плагином
func handleGetNotifier() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/notifier/notifier.html",
		),
	)

	http.HandleFunc(
		"GET /notifier/{messenger}",
		func(w http.ResponseWriter, r *http.Request) {

			messenger := r.PathValue("messenger")

			data, err := notifier(
				messenger,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(
				w,
				"notifier",
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
