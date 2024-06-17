package db

import (
	"errors"
)

type Service struct {
	Name string
	Desc string
}

// Создать таблицу сервисов
func (db *DB) CreateServices() error {

	errMsg := "can't create service table: "

	query := `
		CREATE TABLE service(
			name VARCHAR(32) PRIMARY KEY,
			desc VARCHAR(256)
		);
	`

	_, err := db.Exec(
		query,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Удалить таблицу сервисов
func (db *DB) DropServices() error {

	errMsg := "can't drop service table: "

	query := `
		DROP TABLE service;
	`

	_, err := db.Exec(
		query,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Добавить в сервис в таблицу сервисов
func (db *DB) InsertService(
	s Service,
) error {

	errMsg := "can't insert service: "

	query := `
		INSERT INTO service
		VALUES(
			?,
			?
		);
	`

	_, err := db.Exec(
		query,
		s.Name,
		s.Desc,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать все сервисы
func (db *DB) SelectAllServices() (
	[]Service,
	error,
) {

	errMsg := "can't select services: "

	query := `
		SELECT *
		FROM service;
	`

	rows, err := db.Query(
		query,
	)

	if err != nil {
		return nil,
			errors.New(
				errMsg +
					"can't query: " +
					err.Error(),
			)
	}

	defer rows.Close()

	var services []Service

	for rows.Next() {

		var service Service

		err = rows.Scan(
			&service.Name,
			&service.Desc,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		services = append(
			services,
			service,
		)
	}

	return services,
		nil
}

// Выбрать сервис по имени
func (db *DB) SelectServiceByName(
	name string,
) (
	Service,
	error,
) {

	errMsg := "can't select service: "

	query := `
		SELECT *
		FROM service
		WHERE name = ?;
	`

	row := db.QueryRow(
		query,
		name,
	)

	var service Service

	err := row.Scan(
		&service.Name,
		&service.Desc,
	)

	if err != nil {
		return service,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return service,
		nil
}

// Обновить сервис в таблице сервисов
func (db *DB) UpdateService(
	s Service,
) error {

	errMsg := "can't update service: "

	query := `
		UPDATE service
		SET desc = ?
		WHERE name = ?;
	`

	_, err := db.Exec(
		query,
		s.Desc,
		s.Name,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Удалить сервис из таблицы сервисов
func (db *DB) DeleteService(
	name string,
) error {

	errMsg := "can't delete service: "

	query := `
		DELETE FROM service
		WHERE name = ?;
	`

	_, err := db.Exec(
		query,
		name,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}
