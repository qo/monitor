package value

// Зарегистрировать все обработчики
// для значений
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для получения
	// страницы со значениями
	err = handleGetValues()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения
	// формы со значением
	err = handleGetValue()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для редактирования значения
	err = handleGetEdit()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для создания значения
	err = handleGetNew()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для редактирования значения
	err = handlePut()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для создания значения
	err = handlePost()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления значения
	err = handleDelete()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// с метриками указанного сервиса
	err = handleGetMetricsByService()
	if err != nil {
		return err
	}

	return nil
}
