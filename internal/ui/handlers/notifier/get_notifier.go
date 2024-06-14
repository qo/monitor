package notifier

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы с публикующим плагином
func handleGetNotifier() error {

	// Получить шаблон
	// для формы с публикующим плагином
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/notifier/notifier.html",
		),
	)

	http.HandleFunc(
		"GET /notifier/{messenger}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			messenger := r.PathValue("messenger")

			// Получить публикующий плагин
			data, err := notifier(
				messenger,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы с публикующим плагином
			err = tmpl.ExecuteTemplate(
				w,
				"notifier",
				data,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)

	return nil
}
