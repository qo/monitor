package daisy

import (
	"net/http"
	"os"
)

// Зарегистрировать обработчик для
// получения скрипта daisyUI
func handleGetJs() error {

	// Имя раздаваемого файла
	name := "./internal/ui/external/daisy.min.js"

	// Открыть файл
	_, err := os.Open(name)
	if err != nil {
		return err
	}

	http.HandleFunc(
		"GET /daisy.min.js",
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
