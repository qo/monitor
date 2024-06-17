package trigger

import (
	"log"
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
			"./internal/ui/templates/trigger/user.html",
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
			err = endpointTmpl.ExecuteTemplate(
				w,
				"user",
				data,
			)
			if err != nil {
				log.Fatal(err)
			}
		},
	)

	return nil
}
