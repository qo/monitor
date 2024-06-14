package daisy

// Зарегистрировать все обработчики
// для daisyUI
func Handle() error {
	var err error

	// Зарегистрировать обработчик
	// для скрипта для daisyUI
	err = handleGetJs()
	if err != nil {
		return err
	}

	// Зарегистрировать обработчик
	// для стилей для daisyUI
	return handleGetCss()
}
