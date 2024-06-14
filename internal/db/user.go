package db

import "errors"

type User struct {
	Username string
}

// Создать таблицу пользователей
func (db *DB) CreateUsers() error {

	errMsg := "can't create user table: "

	query := `
		CREATE TABLE user(
			username VARCHAR(32) PRIMARY KEY
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

// Удалить таблицу пользователей
func (db *DB) DropUsers() error {

	errMsg := "can't drop user table: "

	query := `
		DROP TABLE user;
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

// Добавить пользователя в таблицу пользователей
func (db *DB) InsertUser(
	u User,
) error {

	errMsg := "can't insert user: "

	query := `
		INSERT INTO user
		VALUES(
			?
		);
	`

	_, err := db.Exec(
		query,
		u.Username,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать всех пользователей
func (db *DB) SelectAllUsers() (
	[]User,
	error,
) {

	errMsg := "can't select users: "

	query := `
		SELECT *
		FROM user;
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

	var users []User

	for rows.Next() {

		var user User

		err = rows.Scan(
			&user.Username,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		users = append(
			users,
			user,
		)
	}

	return users,
		nil
}

// Выбрать пользователя по имени
func (db *DB) SelectUserByUsername(
	username string,
) (
	User,
	error,
) {

	errMsg := "can't select user: "

	query := `
		SELECT *
		FROM user
		WHERE username = ?;
	`

	row := db.QueryRow(
		query,
		username,
	)

	var user User

	err := row.Scan(
		&user.Username,
	)

	if err != nil {
		return user,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return user,
		nil
}

// Удалить пользователя из таблицы пользователей
func (db *DB) DeleteUser(
	username string,
) error {

	errMsg := "can't delete user: "

	query := `
		DELETE FROM user
		WHERE username = ?;
	`

	_, err := db.Exec(
		query,
		username,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать пользователя по мессенджеру и идентификатору
func (db *DB) SelectUserByMessengerAndId(
	messenger string,
	id string,
) (
	User,
	error,
) {

	errMsg := "can't select user by messenger and id: "

	query := `
		SELECT username FROM endpoint
		WHERE messenger = ?
		AND id = ?;
	`

	row := db.QueryRow(
		query,
		messenger,
		id,
	)

	var user User

	err := row.Scan(
		&user.Username,
	)

	if err != nil {
		return user,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return user,
		nil
}
