package endpoint

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с эндпоинтом
func handleGetEndpoint() error {

	// Шаблон формы с эндпоинтом
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/endpoint.html",
		),
	)

	http.HandleFunc(
		"GET /endpoint/{messenger}/{user}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			messenger := r.PathValue("messenger")
			user := r.PathValue("user")

			// Получить эндпоинт
			data, err := endpoint(
				messenger,
				user,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон формы с эндпоинтом
			err = tmpl.ExecuteTemplate(
				w,
				"endpoint",
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
