package trigger

// Зарегистрировать все обработчики
// для триггера
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для получения страницы с триггерами
	err = handleGetTriggers()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы с триггером
	err = handleGetTrigger()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для создания триггера
	err = handleGetNew()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для создания триггера
	err = handlePost()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления триггера
	err = handleDelete()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы с метриками
	// данного сервиса
	err = handleGetMetricsByService()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы со значениями
	// данной метрики
	err = handleGetValuesByMetric()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы с эндпоинтами
	// данного мессенджера
	err = handleGetEndpointsByMessenger()
	if err != nil {
		return err
	}

	return nil
}
