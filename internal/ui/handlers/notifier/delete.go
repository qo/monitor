package notifier

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления публикующего плагина
func handleDelete() error {

	// Получить шаблон ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /notifiers/remove/{messenger}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			messenger := r.PathValue("messenger")

			// Удалить публикующий плагин
			err := remove(
				messenger,
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
