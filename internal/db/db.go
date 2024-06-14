package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/qo/monitor/internal/config"
)

type DB struct {
	*sql.DB
}

// Открыть БД
func Open() (*DB, func(), error) {

	errMsg := "can't open db: "

	// Загрузить конфигурацию
	config, err := config.Load()
	if err != nil {
		return nil,
			nil,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	// Строка для открытия БД
	dataSourceName := fmt.Sprintf(
		"file:%s?_foreign_keys=on",
		config.DatabasePath,
	)

	// Открыть БД
	_db, err := sql.Open(
		"sqlite3",
		dataSourceName,
	)
	if err != nil {
		return nil,
			nil,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	db := &DB{_db}
	closer := func() {
		db.Close()
	}

	return db,
		closer,
		nil
}

// Удалить все таблицы
func (db *DB) DropTables() error {

	errMsg := "can't drop tables: "

	var err error

	// Удалить таблицу триггеров
	err = db.DropTriggers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить таблицу подписывающих плагинов
	err = db.DropPollers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить таблицу задач
	err = db.DropTasks()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить таблицу значений
	err = db.DropValues()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить таблицу метрик
	err = db.DropMetrics()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить таблицу сервисов
	err = db.DropServices()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить таблицу публикующих плагинов
	err = db.DropNotifiers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить таблицу эндпоинтов
	err = db.DropEndpoints()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить таблицу мессенджеров
	err = db.DropMessengers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить таблицу пользователей
	err = db.DropUsers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Создать все таблицы
func (db *DB) CreateTables() error {

	errMsg := "can't create tables: "

	var err error

	// Создать таблицу сервисов
	err = db.CreateServices()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Создать таблицу метрик
	err = db.CreateMetrics()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Создать таблицу значений
	err = db.CreateValues()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Создать таблицу подписывающих плагинов
	err = db.CreatePollers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Создать таблицу мессенджеров
	err = db.CreateMessengers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Создать таблицу эндпоинтов
	err = db.CreateEndpoints()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Создать таблицу публикующих плагинов
	err = db.CreateNotifiers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Создать таблицу триггеров
	err = db.CreateTriggers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Создать таблицу пользователей
	err = db.CreateUsers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Создать таблицу задач
	err = db.CreateTasks()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}
