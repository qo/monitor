package db

import "errors"

type Notifier struct {
	Messenger string
}

// Создать таблицу публикующих плагинов
func (db *DB) CreateNotifiers() error {

	errMsg := "can't create notifier table: "

	query := `
		CREATE TABLE notifier(
			messenger VARCHAR(32) NOT NULL,
			FOREIGN KEY (messenger)
			REFERENCES messenger(name),
			PRIMARY KEY (messenger)
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

// Удалить таблицу публикующих плагинов
func (db *DB) DropNotifiers() error {

	errMsg := "can't drop notifier table: "

	query := `
		DROP TABLE notifier;
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

// Добавить публикующий плагин в таблицу публикующих плагинов
func (db *DB) InsertNotifier(
	n Notifier,
) error {

	errMsg := "can't insert notifier: "

	query := `
		INSERT INTO notifier
		VALUES(
			?
		);
	`

	_, err := db.Exec(
		query,
		n.Messenger,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать все публикующие плагины
func (db *DB) SelectAllNotifiers() (
	[]Notifier,
	error,
) {

	errMsg := "can't select all notifiers: "

	query := `
		SELECT *
		FROM notifier;
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

	var notifiers []Notifier

	for rows.Next() {

		var notifier Notifier

		err = rows.Scan(
			&notifier.Messenger,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		notifiers = append(
			notifiers,
			notifier,
		)
	}

	return notifiers, nil
}

// Выбрать публикующий плагин по мессенджеру
func (db *DB) SelectNotifierByMessenger(
	messenger string,
) (
	Notifier,
	error,
) {

	errMsg := "can't select notifier: "

	query := `
		SELECT *
		FROM notifier
		WHERE messenger = ?;
	`

	row := db.QueryRow(
		query,
		messenger,
	)

	var notifier Notifier

	err := row.Scan(
		&notifier.Messenger,
	)

	if err != nil {
		return notifier,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return notifier,
		nil
}

// Посчитать количество публикующих плагинов
func (db *DB) CountAllNotifiers() (
	int,
	error,
) {

	errMsg := "can't count all notifiers: "

	query := `
		SELECT COUNT(*)
		FROM notifier;
	`

	row := db.QueryRow(query)

	var count int

	err := row.Scan(&count)

	if err != nil {
		return 0,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return count, nil
}

// Удалить публикующий плагин
// из таблицы публикующих плагинов
func (db *DB) DeleteNotifier(
	messenger string,
) error {

	errMsg := "can't delete notifier: "

	query := `
		DELETE FROM notifier
		WHERE messenger = ?;
	`

	_, err := db.Exec(
		query,
		messenger,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}
