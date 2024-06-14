package endpoint

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с эндпоинтом
func handleGetEndpoint() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/endpoint.html",
		),
	)

	http.HandleFunc(
		"GET /endpoint/{messenger}/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			messenger := r.PathValue("messenger")
			id := r.PathValue("id")
			data, err := endpoint(
				messenger,
				id,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
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
