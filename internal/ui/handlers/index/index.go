package index

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// стартовой страницы
func handleGet() error {

	// Шаблон стартовой страницы
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/index/index.html",
		),
	)

	http.HandleFunc(
		"GET /",
		func(w http.ResponseWriter, r *http.Request) {

			// Отобразить шаблон
			err := tmpl.Execute(
				w,
				nil,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)

	return nil
}
