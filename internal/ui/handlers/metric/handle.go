package metric

// Зарегистрировать все обработчики
// для метрики
func Handle() error {

	var err error

	err = handleGetMetrics()
	if err != nil {
		return err
	}

	err = handleGetMetric()
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
