package service

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования сервиса
func handlePut() error {

	// Получить шаблон
	// для формы с сервисом
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/service.html",
		),
	)

	http.HandleFunc(
		"PUT /service/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			var service db.Service

			// Получить параметры формы
			name := r.FormValue("name")
			desc := r.FormValue("description")

			// Заполнить данные сервиса
			service.Name = name
			service.Desc = desc

			// Обновить сервис
			err := update(
				service,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// с формой с сервисом
			tmpl.ExecuteTemplate(
				w,
				"service",
				service,
			)
		},
	)

	return nil
}
