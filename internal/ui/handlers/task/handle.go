package task

// Зарегистрировать все обработчики
// для задач
func Handle() error {

	var err error

	err = handleGetTasks()
	if err != nil {
		return err
	}

	err = handleGetTask()
	if err != nil {
		return err
	}

	err = handleGetEdit()
	if err != nil {
		return err
	}

	err = handlePut()
	if err != nil {
		return err
	}

	err = handleDelete()
	if err != nil {
		return err
	}

	return nil
}
