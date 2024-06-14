package endpoint

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания эндпоинта
func handlePost() error {

	// Шаблон формы с эндпоинтом
	endpointTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/endpoint.html",
		),
	)

	// Шаблон ошибки 500
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

			// Получить параметры из формы
			messenger := r.FormValue("messenger")
			id := r.FormValue("id")
			desc := r.FormValue("description")

			// Заполнить эндпоинт
			endpoint.Messenger = messenger
			endpoint.Id = id
			endpoint.Desc = desc

			// Добавить эндпоинт
			err := insert(
				endpoint,
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

			// Заполнить шаблон формы с эндпоинтом
			endpointTmpl.ExecuteTemplate(
				w,
				"endpoint",
				endpoint,
			)
		},
	)

	return nil
}
