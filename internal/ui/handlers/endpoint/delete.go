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
		"DELETE /endpoints/remove/{messenger}/{id}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			messenger := r.PathValue("messenger")
			id := r.PathValue("id")

			// Удалить эндпоинт
			err := remove(
				messenger,
				id,
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
