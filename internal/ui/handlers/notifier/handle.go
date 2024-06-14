package notifier

// Зарегистрировать все обработчики
// для публикующих плагинов
func Handle() error {

	var err error

	err = handleGetNotifiers()
	if err != nil {
		return err
	}

	err = handleGetNotifier()
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
