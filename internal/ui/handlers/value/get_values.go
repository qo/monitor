package value

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы со значениями
func handleGetValues() error {

	// Получить шаблон
	// страницы со значениями
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/value/values.html",
			"./internal/ui/templates/value/value.html",
			"./internal/ui/templates/value/new.html",
		),
	)

	http.HandleFunc(
		"GET /values",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить все значения
			data, err := values()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполниь шаблон
			// страницы со значениями
			err = tmpl.Execute(
				w,
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
