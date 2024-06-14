package endpoint

import (
	"net/http"
	"text/template"

	"github.com/qo/monitor/internal/db"
)

// Зарегистрировать обработчик для
// редактирования эндпоинта
func handlePut() error {

	// Шаблон для формы с эндпоинтом
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/endpoint/endpoint.html",
		),
	)

	http.HandleFunc(
		"PUT /endpoint/{messenger}/{name}",
		func(w http.ResponseWriter, r *http.Request) {

			var endpoint db.Endpoint

			// Получить параметры формы
			messenger := r.FormValue("messenger")
			id := r.FormValue("id")
			desc := r.FormValue("description")

			// Заполнить эндпоинт данными
			endpoint.Messenger = messenger
			endpoint.Id = id
			endpoint.Desc = desc

			// Обновить эндпоинт
			err := update(
				endpoint,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон формы с эндпоинтом
			tmpl.ExecuteTemplate(
				w,
				"endpoint",
				endpoint,
			)
		},
	)

	return nil
}
