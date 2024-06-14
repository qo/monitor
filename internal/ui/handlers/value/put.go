package value

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования значения
func handlePut() error {

	// Получить шаблон
	// для формы со значением
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/value.html",
		),
	)

	http.HandleFunc(
		"PUT /value/{service}/{metric}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			var value db.Value

			// Получить параметры формы
			service := r.FormValue("service")
			metric := r.FormValue("metric")
			name := r.FormValue("name")
			desc := r.FormValue("description")
			faultyString := r.FormValue("faulty")

			// Преобразовать строку в bool
			faulty := faultyString != ""

			// Заполнить данные о значении
			value.Service = service
			value.Metric = metric
			value.Name = name
			value.Desc = desc
			value.Faulty = faulty

			// Обновить значение
			err := update(
				value,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для формы со значением
			tmpl.ExecuteTemplate(
				w,
				"value",
				value,
			)
		},
	)

	return nil
}
