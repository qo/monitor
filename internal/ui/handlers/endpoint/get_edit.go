package endpoint

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы редактирования эндпоинта
func handleGetEdit() error {

	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/edit.html",
		),
	)

	notfoundTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/notfound.html",
		),
	)

	http.HandleFunc(
		"GET /endpoints/edit/{messenger}/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			messenger := r.PathValue("messenger")
			name := r.PathValue("name")
			data, err := endpoint(
				messenger,
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
