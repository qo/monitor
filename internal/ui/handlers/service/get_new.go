package service

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для создания нового сервиса
func handleGetNew() error {

	// Получить шаблон формы
	// для создания нового сервиса
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/service/new.html",
		),
	)

	http.HandleFunc(
		"GET /services/new",
		func(w http.ResponseWriter, r *http.Request) {
			// Заполнить шаблон формы
			// для создания нового сервиса
			tmpl.ExecuteTemplate(
				w,
				"new",
				nil,
			)
		},
	)

	return nil
}
