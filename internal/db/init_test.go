package db_test

import (
	"errors"
	"log"
	"os"
	"path"
	"runtime"
)

// Функция, которая вызывается при импорте
// данного пакета.
func init() {

	// Сменить директорию на два уровня вверх
	// относительно текущей,
	// чтобы поместить тестирующий код
	// в контекст корня проекта
	// https://stackoverflow.com/a/60258660
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		log.Fatal(
			errors.New(
				"can't call init function for db tests: " +
					err.Error(),
			),
		)
	}
}
