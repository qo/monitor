package value

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы со значением
func handleGetValue() error {

	// Получить шаблон
	// формы со значением
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/value.html",
		),
	)

	http.HandleFunc(
		"GET /values/{service}/{metric}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить значения query-параметров
			service := r.PathValue("service")
			metric := r.PathValue("metric")
			name := r.PathValue("name")

			// Получить значение
			data, err := value(
				service,
				metric,
				name,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// формы со значением
			err = tmpl.ExecuteTemplate(
				w,
				"value",
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
