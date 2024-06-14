package value

import (
	"log"
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы со значениями
func handleGetValues() error {

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
			data, err := values()
			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
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
