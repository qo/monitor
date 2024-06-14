package poller

// Зарегистрировать все обработчики
// для опрашивающего плагина
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для получения
	// страницы опрашивающих плагинов
	err = handleGetPollers()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения
	// формы с опрашивающим плагином
	err = handleGetPoller()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для создания опрашивающего плагина
	err = handleGetNew()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для создания опрашивающего плагина
	err = handlePost()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления опрашивающего плагина
	err = handleDelete()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения метрик сервиса
	err = handleGetMetricsByService()
	if err != nil {
		return err
	}

	return nil
}
