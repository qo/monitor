package poller

import (
	"log"
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с опрашивающими плагинами
func handleGetPollers() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/poller/pollers.html",
			"./internal/ui/templates/poller/poller.html",
			"./internal/ui/templates/poller/new.html",
		),
	)

	http.HandleFunc(
		"GET /pollers",
		func(w http.ResponseWriter, r *http.Request) {

			data, err := pollers()
			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(
				w,
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
