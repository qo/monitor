package task

// Зарегистрировать все обработчики
// для задач
func Handle() error {

	var err error

	// Зарегистрировать обработчик
	// для получения страницы задач
	err = handleGetTasks()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для получения формы с задачей
	err = handleGetTask()
	if err != nil {
		return err
	}

	return nil
}
