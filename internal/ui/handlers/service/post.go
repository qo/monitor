package service

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания сервиса
func handlePost() error {

	// Получить шаблон
	// для формы с сервисом
	serviceTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/service.html",
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
		"POST /service",
		func(w http.ResponseWriter, r *http.Request) {

			var service db.Service

			// Получить значения query-параметров
			name := r.FormValue("name")
			desc := r.FormValue("description")

			// Заполнить данные сервиса
			service.Name = name
			service.Desc = desc

			// Добавить сервис
			err := insert(
				service,
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
			// для формы с сервисом
			serviceTmpl.ExecuteTemplate(
				w,
				"service",
				service,
			)
		},
	)

	return nil
}
