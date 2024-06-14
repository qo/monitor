package messenger

// Зарегистрировать обработчики
// для мессенджера
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для получения страницы мессенджеров
	err = handleGetMessengers()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы с мессенджером
	err = handleGetMessenger()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для редактирования мессенджера
	err = handleGetEdit()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для создания мессенджера
	err = handleGetNew()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для редактирования мессенджера
	err = handlePut()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для создания мессенджера
	err = handlePost()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления мессенджера
	err = handleDelete()
	if err != nil {
		return err
	}

	return nil
}
