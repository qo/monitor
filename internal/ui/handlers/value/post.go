package value

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// создания значения
func handlePost() error {

	// Получить шаблон
	// для формы со значением
	valueTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/value.html",
		),
	)

	// Получить шаблон для ошибки 500
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

			// Получить параметры формы
			service := r.FormValue("service")
			metric := r.FormValue("metric")
			name := r.FormValue("name")
			desc := r.FormValue("description")
			faultyString := r.FormValue("faulty")

			// Преобразовать строку к bool
			faulty := faultyString != ""

			// Заполнить данные значения
			value.Service = service
			value.Metric = metric
			value.Name = name
			value.Desc = desc
			value.Faulty = faulty

			// Добавить значение
			err := insert(
				value,
			)
			if err != nil {
				// Заполнить шаблон для ошибки 500
				internalErrorTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// для формы со значением
			valueTmpl.ExecuteTemplate(
				w,
				"value",
				value,
			)
		},
	)

	return nil
}
