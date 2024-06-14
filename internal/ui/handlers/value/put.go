package value

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования значения
func handlePut() error {

	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/value.html",
		),
	)

	http.HandleFunc(
		"PUT /value/{service}/{metric}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			var value db.Value

			service := r.FormValue("service")
			value.Service = service

			metric := r.FormValue("metric")
			value.Metric = metric

			name := r.FormValue("name")
			value.Name = name

			desc := r.FormValue("description")
			value.Desc = desc

			faultyString := r.FormValue("faulty")
			faulty := faultyString != ""
			value.Faulty = faulty

			err := update(value)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tmpl.ExecuteTemplate(
				w,
				"value",
				value,
			)
		},
	)

	return nil
}
