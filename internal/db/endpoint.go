package db

import "errors"

type Endpoint struct {
	Messenger string
	User      string
	Id        string
	Desc      string
}

// Создать таблицу эндпоинтов
func (db *DB) CreateEndpoints() error {

	errMsg := "can't create endpoint table: "

	query := `
		CREATE TABLE endpoint(
			messenger VARCHAR(32) NOT NULL,
			user VARCHAR(32) NOT NULL,
			id VARCHAR(32) NOT NULL,
			desc VARCHAR(256),
			FOREIGN KEY (messenger)
			REFERENCES messenger(name),
			FOREIGN KEY (user)
			REFERENCES user(username),
			PRIMARY KEY (messenger, user),
			UNIQUE (messenger, id),
			CHECK(id <> '')
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

// Удалить таблицу эндпоинтов
func (db *DB) DropEndpoints() error {

	errMsg := "can't drop endpoint table: "

	query := `
		DROP TABLE endpoint;
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

// Добавить эндпоинт в таблицу эндпоинтов
func (db *DB) InsertEndpoint(
	e Endpoint,
) error {

	errMsg := "can't insert endpoint: "

	query := `
		INSERT INTO endpoint
		VALUES(
			?,
			?,
			?,
			?
		);
	`

	_, err := db.Exec(
		query,
		e.Messenger,
		e.User,
		e.Id,
		e.Desc,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать все эндпоинты
func (db *DB) SelectAllEndpoints() (
	[]Endpoint,
	error,
) {

	errMsg := "can't select all endpoints: "

	query := `
		SELECT *
		FROM endpoint;
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

	var endpoints []Endpoint

	for rows.Next() {

		var endpoint Endpoint

		err = rows.Scan(
			&endpoint.Messenger,
			&endpoint.User,
			&endpoint.Id,
			&endpoint.Desc,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		endpoints = append(
			endpoints,
			endpoint,
		)
	}

	return endpoints,
		nil
}

// Выбрать эндпоинт по мессенджеру и пользователю
func (db *DB) SelectEndpointByMessengerAndUser(
	messenger string,
	user string,
) (
	Endpoint,
	error,
) {

	errMsg := "can't select endpoint: "

	query := `
		SELECT *
		FROM endpoint
		WHERE messenger = ?
		AND user = ?;
	`

	row := db.QueryRow(
		query,
		messenger,
		user,
	)

	var endpoint Endpoint

	err := row.Scan(
		&endpoint.Messenger,
		&endpoint.User,
		&endpoint.Id,
		&endpoint.Desc,
	)

	if err != nil {
		return endpoint,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return endpoint,
		nil
}

// Выбрать эндпоинт по мессенджеру и идентификатору
func (db *DB) SelectEndpointByMessengerAndId(
	messenger string,
	id string,
) (
	Endpoint,
	error,
) {

	errMsg := "can't select endpoint: "

	query := `
		SELECT *
		FROM endpoint
		WHERE messenger = ?
		AND id = ?;
	`

	row := db.QueryRow(
		query,
		messenger,
		id,
	)

	var endpoint Endpoint

	err := row.Scan(
		&endpoint.Messenger,
		&endpoint.User,
		&endpoint.Id,
		&endpoint.Desc,
	)

	if err != nil {
		return endpoint,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return endpoint,
		nil
}

// Выбрать эндпоинты по мессенджеру
func (db *DB) SelectEndpointsByMessenger(
	messenger string,
) (
	[]Endpoint,
	error,
) {

	errMsg := "can't select endpoints by messenger: "

	query := `
		SELECT *
		FROM endpoint
		WHERE messenger = ?;
	`

	rows, err := db.Query(
		query,
		messenger,
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

	var endpoints []Endpoint

	for rows.Next() {

		var endpoint Endpoint

		err = rows.Scan(
			&endpoint.Messenger,
			&endpoint.User,
			&endpoint.Id,
			&endpoint.Desc,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		endpoints = append(
			endpoints,
			endpoint,
		)
	}

	return endpoints,
		nil
}

// Обновить эндпоинт из таблицы эндпоинтов
func (db *DB) UpdateEndpoint(
	e Endpoint,
) error {

	errMsg := "can't update endpoint: "

	query := `
		UPDATE endpoint
		SET id = ?,
		desc = ?
		WHERE messenger = ?
		AND user = ?;
	`

	_, err := db.Exec(
		query,
		e.Id,
		e.Desc,
		e.Messenger,
		e.User,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Удалить эндпоинт из таблицы эндпоинтов
func (db *DB) DeleteEndpoint(
	messenger string,
	user string,
) error {

	errMsg := "can't delete endpoint: "

	query := `
		DELETE FROM endpoint
		WHERE messenger = ?
		AND user = ?;
	`

	_, err := db.Exec(
		query,
		messenger,
		user,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}
