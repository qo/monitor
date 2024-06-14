package service

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования сервиса
func handleGetEdit() error {

	// Получить шаблон формы
	// для редактирования сервиса
	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/edit.html",
		),
	)

	// Получить шаблон ошибки 404
	notfoundTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/notfound.html",
		),
	)

	http.HandleFunc(
		"GET /services/edit/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			name := r.PathValue("name")

			// Получить сервис
			data, err := service(
				name,
			)
			if err != nil {
				// Заполнить шаблон ошибки 404
				notfoundTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон формы
			// для редактирования сервиса
			editTmpl.ExecuteTemplate(
				w,
				"edit",
				data,
			)
		},
	)

	return nil
}
