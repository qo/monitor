package db

import "errors"

type Messenger struct {
	Name string
	Desc string
}

// Создать таблицу мессенджеров
func (db *DB) CreateMessengers() error {

	errMsg := "can't create messenger table: "

	query := `
		CREATE TABLE messenger(
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

// Удалить таблицу мессенджеров
func (db *DB) DropMessengers() error {

	errMsg := "can't drop messenger table: "

	query := `
		DROP TABLE messenger;
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

// Добавить мессенджер в таблицу мессенджеров
func (db *DB) InsertMessenger(
	m Messenger,
) error {

	errMsg := "can't insert messenger: "

	query := `
		INSERT INTO messenger
		VALUES(
			?,
			?
		);
	`

	_, err := db.Exec(
		query,
		m.Name,
		m.Desc,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать все мессенджеры из таблицы мессенджеров
func (db *DB) SelectAllMessengers() (
	[]Messenger,
	error,
) {

	errMsg := "can't select all messengers: "

	query := `
		SELECT *
		FROM messenger;
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

	var messengers []Messenger

	for rows.Next() {

		var messenger Messenger

		err = rows.Scan(
			&messenger.Name,
			&messenger.Desc,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		messengers = append(
			messengers,
			messenger,
		)
	}

	return messengers,
		nil
}

// Выбрать мессенджер по имени
func (db *DB) SelectMessengerByName(
	name string,
) (
	Messenger,
	error,
) {

	errMsg := "can't select messenger: "

	query := `
		SELECT *
		FROM messenger
		WHERE name = ?;
	`

	row := db.QueryRow(
		query,
		name,
	)

	var messenger Messenger

	err := row.Scan(
		&messenger.Name,
		&messenger.Desc,
	)

	if err != nil {
		return messenger,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return messenger,
		nil
}

// Обновить мессенджер в таблице мессенджеров
func (db *DB) UpdateMessenger(
	m Messenger,
) error {

	errMsg := "can't update messenger: "

	query := `
		UPDATE messenger
		SET desc = ?
		WHERE name = ?;
	`

	_, err := db.Exec(
		query,
		m.Desc,
		m.Name,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Удалить мессенджеров из таблицы мессенджеров
func (db *DB) DeleteMessenger(
	name string,
) error {

	errMsg := "can't delete messenger: "

	query := `
		DELETE FROM messenger
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
