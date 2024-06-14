package value

// Зарегистрировать все обработчики
// для значения
func Handle() error {

	var err error

	err = handleGetValues()
	if err != nil {
		return err
	}

	err = handleGetValue()
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

	err = handleGetMetricsByService()
	if err != nil {
		return err
	}

	return nil
}
