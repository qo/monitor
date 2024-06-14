package htmx

import (
	"net/http"
	"os"
)

// Зарегистрировать обработчик для
// получения скрипта htmx
func handleGet() error {
	name := "./internal/ui/external/htmx.min.js"
	_, err := os.Open(name)
	if err != nil {
		return err
	}
	http.HandleFunc(
		"GET /htmx.min.js",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(
				w,
				r,
				name,
			)
		},
	)
	return nil
}
