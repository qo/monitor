package value

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания значения
func handlePost() error {

	valueTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/value.html",
		),
	)

	internalErrorTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/internal.html",
		),
	)

	http.HandleFunc(
		"POST /value",
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

			err := insert(value)
			if err != nil {
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			valueTmpl.ExecuteTemplate(
				w,
				"value",
				value,
			)
		},
	)

	return nil
}
