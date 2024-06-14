package daisy

import (
	"net/http"
	"os"
)

// Зарегистрировать обработчик для
// получения скрипта daisyUI
func handleGetJs() error {
	name := "./internal/ui/external/daisy.min.js"
	_, err := os.Open(name)
	if err != nil {
		return err
	}
	http.HandleFunc(
		"GET /daisy.min.js",
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
