package metric

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// удаления метрики
func handleDelete() error {

	// Получить шаблон ошибки 500
	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"DELETE /metrics/remove/{service}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.PathValue("service")
			name := r.PathValue("name")

			// Удалить метрику
			err := remove(
				service,
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
