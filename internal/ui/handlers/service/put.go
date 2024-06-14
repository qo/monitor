package service

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования сервиса
func handlePut() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/service.html",
		),
	)

	http.HandleFunc(
		"PUT /service/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			var service db.Service

			name := r.FormValue("name")
			service.Name = name

			desc := r.FormValue("description")
			service.Desc = desc

			err := update(service)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tmpl.ExecuteTemplate(
				w,
				"service",
				service,
			)
		},
	)

	return nil
}
