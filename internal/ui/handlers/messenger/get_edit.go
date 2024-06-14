package messenger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования мессенджера
func handleGetEdit() error {

	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/messenger/edit.html",
		),
	)

	notfoundTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/notfound.html",
		),
	)

	http.HandleFunc(
		"GET /messengers/edit/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			name := r.PathValue("name")
			data, err := messenger(
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
