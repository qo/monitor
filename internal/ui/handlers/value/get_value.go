package value

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы со значением
func handleGetValue() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/value.html",
		),
	)

	http.HandleFunc(
		"GET /values/{service}/{metric}/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			service := r.PathValue("service")
			metric := r.PathValue("metric")
			name := r.PathValue("name")
			data, err := value(
				service,
				metric,
				name,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteTemplate(
				w,
				"value",
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
