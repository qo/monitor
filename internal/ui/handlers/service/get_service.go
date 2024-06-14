package service

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с сервисом
func handleGetService() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/service.html",
		),
	)

	http.HandleFunc(
		"GET /service/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			name := r.PathValue("name")
			data, err := service(
				name,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteTemplate(
				w,
				"service",
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
