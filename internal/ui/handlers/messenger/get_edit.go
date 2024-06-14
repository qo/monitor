package messenger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования мессенджера
func handleGetEdit() error {

	// Шаблон формы
	// для редактирования эндпоинта
	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/messenger/edit.html",
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
		"GET /messengers/edit/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			name := r.PathValue("name")

			// Получить мессенджер
			data, err := messenger(
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

			// Заполнить шаблон
			// формы редактирования мессенджера
			editTmpl.ExecuteTemplate(
				w,
				"edit",
				data,
			)
		},
	)

	return nil
}
