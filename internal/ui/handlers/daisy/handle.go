package daisy

// Зарегистрировать все обработчики
// для daisyUI
func Handle() error {
	var err error

	err = handleGetJs()
	if err != nil {
		return err
	}

	return handleGetCss()
}
