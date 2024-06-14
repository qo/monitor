package trigger

// Зарегистрировать все обработчики
// для триггера
func Handle() error {

	var err error

	err = handleGetTriggers()
	if err != nil {
		return err
	}

	err = handleGetTrigger()
	if err != nil {
		return err
	}

	// err = handleGetEdit()
	// if err != nil {
	// 	return err
	// }

	err = handleGetNew()
	if err != nil {
		return err
	}

	// err = handlePut()
	// if err != nil {
	// 	return err
	// }

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

	err = handleGetValuesByMetric()
	if err != nil {
		return err
	}

	err = handleGetEndpointsByMessenger()
	if err != nil {
		return err
	}

	return nil
}
