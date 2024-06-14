package htmx

import (
	"net/http"
	"os"
)

// Зарегистрировать обработчик для
// получения скрипта htmx
func handleGet() error {

	// Имя раздаваемого файла
	name := "./internal/ui/external/htmx.min.js"

	// Открыть файл
	_, err := os.Open(name)
	if err != nil {
		return err
	}

	http.HandleFunc(
		"GET /htmx.min.js",
		func(w http.ResponseWriter, r *http.Request) {
			// Раздать файл
			http.ServeFile(
				w,
				r,
				name,
			)
		},
	)

	return nil
}
