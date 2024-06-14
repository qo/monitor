package service

// Зарегистрировать все обработчики
// для сервиса
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для получения страницы сервисов
	err = handleGetServices()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы с сервисом
	err = handleGetService()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для редактирования сервиса
	err = handleGetEdit()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для создания сервиса
	err = handleGetNew()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для редактирования сервиса
	err = handlePut()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для создания сервиса
	err = handlePost()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления сервиса
	err = handleDelete()
	if err != nil {
		return err
	}

	return nil
}
