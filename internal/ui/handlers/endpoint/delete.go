package endpoint

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления эндпоинта
func handleDelete() error {

	// Шаблон для сообщения об ошибке
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /endpoints/remove/{messenger}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			messenger := r.PathValue("messenger")
			name := r.PathValue("name")

			// Удалить эндпоинт
			err := remove(
				messenger,
				name,
			)
			if err != nil {
				// Заполнить шаблон с ошибкой
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
