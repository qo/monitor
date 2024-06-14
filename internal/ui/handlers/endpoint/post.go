package endpoint

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания эндпоинта
func handlePost() error {

	endpointTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/endpoint.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /endpoint",
		func(w http.ResponseWriter, r *http.Request) {

			var endpoint db.Endpoint

			messenger := r.FormValue("messenger")
			endpoint.Messenger = messenger

			id := r.FormValue("id")
			endpoint.Id = id

			desc := r.FormValue("description")
			endpoint.Desc = desc

			err := insert(endpoint)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			endpointTmpl.ExecuteTemplate(
				w,
				"endpoint",
				endpoint,
			)
		},
	)

	return nil
}
