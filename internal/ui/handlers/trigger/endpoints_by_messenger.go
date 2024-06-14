package trigger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с эндпоинтами
// указанного мессенджера
func handleGetEndpointsByMessenger() error {

	endpointTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/endpoint.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"GET /endpoints-by-messenger",
		func(w http.ResponseWriter, r *http.Request) {

			messenger := r.FormValue("messenger")

			data, err := endpoints(
				messenger,
			)
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
				data,
			)
		},
	)

	return nil
}
