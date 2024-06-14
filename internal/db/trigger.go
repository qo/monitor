package db

import "errors"

type Trigger struct {
	Service   string
	Metric    string
	Value     string
	Messenger string
	User      string
}

// Создать таблицу триггеров
func (db *DB) CreateTriggers() error {

	errMsg := "can't create trigger table: "

	query := `
		CREATE TABLE trigger(
			service VARCHAR(32) NOT NULL,
			metric VARCHAR(32) NOT NULL,
			value VARCHAR(32) NOT NULL,
			messenger VARCHAR(32) NOT NULL,
			user VARCHAR(32) NOT NULL,
			FOREIGN KEY (
				service,
				metric,
				value
			)
			REFERENCES value(
				service,
				metric,
				name
			),
			FOREIGN KEY (
				messenger,
				user
			)
			REFERENCES endpoint(
				messenger,
				user
			),
			PRIMARY KEY (
				service,
				metric,
				value,
				messenger,
				user
			)
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

// Удалить таблицу триггеров
func (db *DB) DropTriggers() error {

	errMsg := "can't drop trigger table: "

	query := `
		DROP TABLE trigger;
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

// Добавить триггер в таблицу триггеров
func (db *DB) InsertTrigger(
	t Trigger,
) error {

	errMsg := "can't insert trigger: "

	query := `
		INSERT INTO trigger
		VALUES(
			?,
			?,
			?,
			?,
			?
		);
	`

	_, err := db.Exec(
		query,
		t.Service,
		t.Metric,
		t.Value,
		t.Messenger,
		t.User,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать все триггеры
func (db *DB) SelectAllTriggers() (
	[]Trigger,
	error,
) {

	errMsg := "can't select triggers: "

	query := `
		SELECT *
		FROM trigger;
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

	var triggers []Trigger

	for rows.Next() {

		var trigger Trigger

		err = rows.Scan(
			&trigger.Service,
			&trigger.Metric,
			&trigger.Value,
			&trigger.Messenger,
			&trigger.User,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		triggers = append(
			triggers,
			trigger,
		)
	}

	return triggers,
		nil
}

// Выбрать триггер по сервису, метрике, значению, мессенджеру и пользователю
func (db *DB) SelectTriggerByServiceMetricValueMessengerAndUser(
	service string,
	metric string,
	value string,
	messenger string,
	user string,
) (
	Trigger,
	error,
) {

	errMsg := "can't select trigger: "

	query := `
		SELECT *
		FROM trigger
		WHERE service = ?
		AND metric= ?
		AND value = ?
		AND messenger = ?
		AND user = ?;
	`

	row := db.QueryRow(
		query,
		messenger,
	)

	var trigger Trigger

	err := row.Scan(
		&trigger.Service,
		&trigger.Metric,
		&trigger.Value,
		&trigger.Messenger,
		&trigger.User,
	)

	return trigger,
		errors.New(
			errMsg +
				err.Error(),
		)
}

// Удалить триггер из таблицы триггеров
func (db *DB) DeleteTrigger(
	service string,
	metric string,
	value string,
	messenger string,
	user string,
) error {

	errMsg := "can't delete trigger: "

	query := `
		DELETE FROM trigger
		WHERE service = ?
		AND metric = ?
		AND value = ?
		AND messenger = ?
		AND user = ?;
	`

	_, err := db.Exec(
		query,
		service,
		metric,
		value,
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
