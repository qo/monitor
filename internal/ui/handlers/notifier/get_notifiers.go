package notifier

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы публикующих плагинов
func handleGetNotifiers() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/notifier/notifiers.html",
			"./internal/ui/templates/notifier/notifier.html",
			"./internal/ui/templates/notifier/new.html",
		),
	)

	http.HandleFunc(
		"GET /notifiers",
		func(w http.ResponseWriter, r *http.Request) {

			data, err := notifiers()
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
