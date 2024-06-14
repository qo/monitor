package endpoint

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы редактирования эндпоинта
func handleGetEdit() error {

	// Шаблон формы для редактирования эндпоинта
	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/edit.html",
		),
	)

	// Шаблон ошибки 404
	notfoundTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/notfound.html",
		),
	)

	http.HandleFunc(
		"GET /endpoints/edit/{messenger}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			messenger := r.PathValue("messenger")
			name := r.PathValue("name")

			// Получить эндпоинт
			data, err := endpoint(
				messenger,
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

			// Заполнить шаблон формы для редактирования эндпоинта
			editTmpl.ExecuteTemplate(
				w,
				"edit",
				data,
			)
		},
	)

	return nil
}
