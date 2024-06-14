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

	// Зарегистрировать обработчик
	// для получения формы
	// для редактирования задачи
	err = handleGetEdit()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для редактирования задачи
	err = handlePut()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для удаления задачи
	err = handleDelete()
	if err != nil {
		return err
	}

	return nil
}
