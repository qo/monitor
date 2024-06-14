package service

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с сервисом
func handleGetService() error {

	// Получить шаблон
	// для формы с сервисом
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/service.html",
		),
	)

	http.HandleFunc(
		"GET /service/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			name := r.PathValue("name")

			// Получить сервис
			data, err := service(
				name,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы с сервисом
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
