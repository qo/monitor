package metric

// Зарегистрировать все обработчики
// для метрики
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для получения страницы метрик
	err = handleGetMetrics()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для формы с метрикой
	err = handleGetMetric()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для формы
	// для редактирования метрики
	err = handleGetEdit()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для формы
	// для создания метрики
	err = handleGetNew()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для редактирования метрики
	err = handlePut()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для создания метрики
	err = handlePost()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления метрики
	err = handleDelete()
	if err != nil {
		return err
	}

	return nil
}
