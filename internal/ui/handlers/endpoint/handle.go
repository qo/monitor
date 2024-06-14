package endpoint

// Зарегистрировать все обработчики для эндпоинтов
func Handle() error {

	var err error

	err = handleGetEndpoints()
	if err != nil {
		return err
	}

	err = handleGetEndpoint()
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
