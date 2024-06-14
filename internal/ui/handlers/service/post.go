package service

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания сервиса
func handlePost() error {

	serviceTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/service.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /service",
		func(w http.ResponseWriter, r *http.Request) {

			var service db.Service

			name := r.FormValue("name")
			service.Name = name

			desc := r.FormValue("description")
			service.Desc = desc

			err := insert(service)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			serviceTmpl.ExecuteTemplate(
				w,
				"service",
				service,
			)
		},
	)

	return nil
}
