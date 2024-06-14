package service

// Зарегистрировать все обработчики
// для сервиса
func Handle() error {

	var err error

	err = handleGetServices()
	if err != nil {
		return err
	}

	err = handleGetService()
	if err != nil {
		return err
	}

	err = handleGetEdit()
	if err != nil {
		return err
	}

	err = handleGetNew()
	if err != nil {
		return err
	}

	err = handlePut()
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
