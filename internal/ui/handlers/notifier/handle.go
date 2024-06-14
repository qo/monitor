package notifier

// Зарегистрировать все обработчики
// для публикующих плагинов
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для получения
	// страницы публикующих плагинов
	err = handleGetNotifiers()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения
	// формы с публикующим плагином
	err = handleGetNotifier()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для создания публикующего плагина
	err = handleGetNew()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для создания публикующего плагина
	err = handlePost()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления публикующего плагина
	err = handleDelete()
	if err != nil {
		return err
	}

	return nil
}
