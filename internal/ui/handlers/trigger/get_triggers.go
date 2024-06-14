package trigger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с триггерами
func handleGetTriggers() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/trigger/triggers.html",
			"./internal/ui/templates/trigger/trigger.html",
			"./internal/ui/templates/trigger/new.html",
		),
	)

	http.HandleFunc(
		"GET /triggers",
		func(w http.ResponseWriter, r *http.Request) {
			data, err := triggers()
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
