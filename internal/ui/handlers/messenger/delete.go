package messenger

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления мессенджера
func handleDelete() error {

	// Шаблон для ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /messengers/remove/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			name := r.PathValue("name")

			// Удалить мессенджер
			err := remove(
				name,
			)
			if err != nil {
				// Заполнить шаблон ошибки 500
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}
		},
	)

	return nil
}
