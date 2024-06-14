package daisy

import (
	"net/http"
	"os"
)

// Зарегистрировать обработчик для
// получения стилей daisyUI
func handleGetCss() error {

	// Имя раздаваемого файла
	name := "./internal/ui/external/daisy.min.css"

	// Открыть файл
	_, err := os.Open(name)
	if err != nil {
		return err
	}

	http.HandleFunc(
		"GET /daisy.min.css",
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
