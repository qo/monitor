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

	config, err := config.Load()
	if err != nil {
		return nil,
			nil,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	dataSourceName := fmt.Sprintf(
		"file:%s?_foreign_keys=on",
		config.DatabasePath,
	)

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

	err = db.DropTriggers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.DropPollers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.DropTasks()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.DropValues()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.DropMetrics()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.DropServices()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.DropNotifiers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.DropEndpoints()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.DropMessengers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

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

	err = db.CreateServices()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.CreateMetrics()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.CreateValues()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.CreatePollers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.CreateMessengers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.CreateEndpoints()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.CreateNotifiers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.CreateTriggers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.CreateUsers()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	err = db.CreateTasks()
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}
