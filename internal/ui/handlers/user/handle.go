package user

// Зарегистрировать все обработчики
// для пользователя
func Handle() error {

	var err error

	err = handleGetUsers()
	if err != nil {
		return err
	}

	err = handleGetUser()
	if err != nil {
		return err
	}

	err = handleGetNew()
	if err != nil {
		return err
	}

	err = handlePost()
	if err != nil {
		return err
	}

	err = handleDelete()
	if err != nil {
		return err
	}

	return nil
}
