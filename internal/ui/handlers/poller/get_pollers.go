package poller

import (
	"net/http"
	"text/template"
)

// Зарегистрировать обработчик для
// получения страницы с опрашивающими плагинами
func handleGetPollers() error {

	// Получить шаблон
	// для страницы с опрашивающими плагинами
	tmpl := template.Must(
		template.ParseFiles(
			"./internal/ui/templates/page.html",
			"./internal/ui/templates/poller/pollers.html",
			"./internal/ui/templates/poller/poller.html",
			"./internal/ui/templates/poller/new.html",
		),
	)

	http.HandleFunc(
		"GET /pollers",
		func(w http.ResponseWriter, r *http.Request) {

			// Получить все опрашивающие плагины
			data, err := pollers()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Заполнить шаблон
			// для страницы с опрашивающими плагинами
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
