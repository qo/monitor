package service

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для создания нового сервиса
func handleGetNew() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/new.html",
		),
	)

	http.HandleFunc(
		"GET /services/new",
		func(w http.ResponseWriter, r *http.Request) {
			if err != nil {
				// TODO: do something
			}
			tmpl.ExecuteTemplate(
				w,
				"new",
				nil,
			)
		},
	)

	return nil
}
