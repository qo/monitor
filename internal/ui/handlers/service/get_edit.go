package service

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования сервиса
func handleGetEdit() error {

	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/edit.html",
		),
	)

	notfoundTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/notfound.html",
		),
	)

	http.HandleFunc(
		"GET /services/edit/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			name := r.PathValue("name")
			data, err := service(
				name,
			)
			if err != nil {
				notfoundTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}
			editTmpl.ExecuteTemplate(
				w,
				"edit",
				data,
			)
		},
	)

	return nil
}
