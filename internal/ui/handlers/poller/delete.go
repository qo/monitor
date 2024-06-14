package poller

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления опрашивающего плагина
func handleDelete() error {

	// Шаблон ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /pollers/remove/{service}/{metric}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.PathValue("service")
			metric := r.PathValue("metric")

			// Удалить подписывающий плагин
			err := remove(
				service,
				metric,
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
