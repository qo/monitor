package user

// Зарегистрировать все обработчики
// для пользователя
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для получения страницы с пользователями
	err = handleGetUsers()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы с пользователем
	err = handleGetUser()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы
	// для создания пользователя
	err = handleGetNew()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для создания пользователя
	err = handlePost()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления пользователя
	err = handleDelete()
	if err != nil {
		return err
	}

	return nil
}
