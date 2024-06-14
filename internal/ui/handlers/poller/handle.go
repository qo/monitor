package poller

// Зарегистрировать все обработчики
// для опрашивающего плагина
func Handle() error {

	var err error

	err = handleGetPollers()
	if err != nil {
		return err
	}

	err = handleGetPoller()
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

	err = handleGetMetricsByService()
	if err != nil {
		return err
	}

	return nil
}
