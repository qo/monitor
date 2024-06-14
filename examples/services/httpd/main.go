package httpd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/qo/monitor/examples/services/httpd/config"
)

// Создать веб-обработчик
func CreateHandler(
	response string,
) func(
	http.ResponseWriter,
	*http.Request,
) {

	return func(
		w http.ResponseWriter,
		req *http.Request,
	) {
		fmt.Fprintf(
			w,
			response,
		)
	}
}

func main() {

	// Загрузить конфигурацию HTTP-сервера
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	handler := CreateHandler(
		cfg.Response,
	)

	// Зарегистрировать обработчик Hello из пакета httpd
	// на все URI
	http.HandleFunc(
		"/",
		handler,
	)

	// Запустить HTTP-сервер
	// на указанных хосте и порте
	http.ListenAndServe(
		fmt.Sprintf(
			"%s:%d",
			cfg.Host,
			cfg.Port,
		),
		nil,
	)
}
