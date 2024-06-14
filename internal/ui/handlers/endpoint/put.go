package endpoint

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования эндпоинта
func handlePut() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/endpoint.html",
		),
	)

	http.HandleFunc(
		"PUT /endpoint/{messenger}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			var endpoint db.Endpoint

			messenger := r.FormValue("messenger")
			endpoint.Messenger = messenger

			id := r.FormValue("id")
			endpoint.Id = id

			desc := r.FormValue("description")
			endpoint.Desc = desc

			err := update(endpoint)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tmpl.ExecuteTemplate(
				w,
				"endpoint",
				endpoint,
			)
		},
	)

	return nil
}
