package trigger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с эндпоинтами
// указанного мессенджера
func handleGetEndpointsByMessenger() error {

	// Получить шаблон
	// для формы с эндпоинтами
	endpointTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/trigger/endpoint.html",
		),
	)

	// Получить шаблон ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"GET /endpoints-by-messenger",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить параметры формы
			messenger := r.FormValue("messenger")

			// Получить эндпоинты данного мессендежера
			data, err := endpoints(
				messenger,
			)
			if err != nil {
				// Заполнить шаблон ошибки 500
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// формы с эндпоинтами
			endpointTmpl.ExecuteTemplate(
				w,
				"endpoint",
				data,
			)
		},
	)

	return nil
}
