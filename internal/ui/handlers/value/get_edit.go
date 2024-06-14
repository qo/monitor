package value

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения формы для редактирования значения
func handleGetEdit() error {

	// Получить шаблон
	// для формы
	// для редактирования значения
	editTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/value/edit.html",
		),
	)

	// Получить шаблон ошибки 404
	notfoundTmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/error/error.html",
			"./internal/ui/templates/error/notfound.html",
		),
	)

	http.HandleFunc(
		"GET /values/edit/{service}/{metric}/{name}",
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
				// Заполнить шаблон ошибки 404
				notfoundTmpl.ExecuteTemplate(
					w,
					"error",
					err.Error(),
				)
				return
			}

			// Заполнить шаблон
			// для формы
			// для редактирования значения
			editTmpl.ExecuteTemplate(
				w,
				"edit",
				data,
			)
		},
	)

	return nil
}
