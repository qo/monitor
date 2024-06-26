package endpoint

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы эндпоинтов
func handleGetEndpoints() error {

	// Шаблон формы страницы эндпоинтов
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/endpoint/endpoints.html",
			"./internal/ui/templates/endpoint/endpoint.html",
			"./internal/ui/templates/endpoint/new.html",
		),
	)

	http.HandleFunc(
		"GET /endpoints",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить все эндпоинты
			data, err := endpoints()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить страницу эндпоинтов
			tmpl.Execute(
				w,
				data,
			)
		},
	)

	return nil
}
