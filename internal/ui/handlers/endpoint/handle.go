package endpoint

// Зарегистрировать все обработчики для эндпоинтов
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для страницы эндпоинтов
	err = handleGetEndpoints()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для шаблона с эндпоинтом
	err = handleGetEndpoint()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для шаблона
	// для формы редактирования эндпоинта
	err = handleGetEdit()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для шаблона
	// для формы создания эндпоинта
	err = handleGetNew()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для редактирования эндпоинта
	err = handlePut()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для создания эндпоинта
	err = handlePost()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления эндпоинта
	err = handleDelete()
	if err != nil {
		return err
	}

	return nil
}
