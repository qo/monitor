package service

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с сервисами
func handleGetServices() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/service/services.html",
			"./internal/ui/templates/service/service.html",
			"./internal/ui/templates/service/new.html",
		),
	)

	http.HandleFunc(
		"GET /services",
		func(w http.ResponseWriter, r *http.Request) {
			data, err := services()
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
