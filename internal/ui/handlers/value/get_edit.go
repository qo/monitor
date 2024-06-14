package value

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования значения
func handleGetEdit() error {

	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/edit.html",
		),
	)

	notfoundTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/notfound.html",
		),
	)

	http.HandleFunc(
		"GET /values/edit/{service}/{metric}/{name}",
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
