package db

import "errors"

type Value struct {
	Service string
	Metric  string
	Name    string
	Desc    string
	Faulty  bool
}

// Создать таблицу значений
func (db *DB) CreateValues() error {

	errMsg := "can't create value table: "

	query := `
		CREATE TABLE value(
			service VARCHAR(32) NOT NULL,
			metric VARCHAR(32) NOT NULL,
			name VARCHAR(32) NOT NULL,
			desc VARCHAR(256),
			faulty INTEGER NOT NULL,
			FOREIGN KEY (
				service,
				metric
			)
			REFERENCES metric(
				service,
				name
			),
			PRIMARY KEY (
				service,
				metric,
				name
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

// Удалить таблицу значений
func (db *DB) DropValues() error {

	errMsg := "can't drop value table: "

	query := `
		DROP TABLE value;
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

// Добавить значение в таблицу значений
func (db *DB) InsertValue(
	v Value,
) error {

	errMsg := "can't insert value: "

	query := `
		INSERT INTO value
		VALUES(
			?,
			?,
			?,
			?,
			?
		);
	`

	var faulty int
	if v.Faulty {
		faulty = 1
	} else {
		faulty = 0
	}

	_, err := db.Exec(
		query,
		v.Service,
		v.Metric,
		v.Name,
		v.Desc,
		faulty,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать все значения
func (db *DB) SelectAllValues() (
	[]Value,
	error,
) {

	errMsg := "can't select values: "

	query := `
		SELECT *
		FROM value;
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

	var values []Value

	for rows.Next() {

		var value Value

		err = rows.Scan(
			&value.Service,
			&value.Metric,
			&value.Name,
			&value.Desc,
			&value.Faulty,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		values = append(
			values,
			value,
		)
	}

	return values,
		nil
}

// Выбрать значение по сервису, метрике и имени
func (db *DB) SelectValueByServiceMetricAndName(
	service string,
	metric string,
	name string,
) (
	Value,
	error,
) {

	errMsg := "can't select value: "

	query := `
		SELECT *
		FROM value
		WHERE service = ?
		AND metric = ?
		AND name = ?;
	`

	row := db.QueryRow(
		query,
		service,
		metric,
		name,
	)

	var value Value

	err := row.Scan(
		&value.Service,
		&value.Metric,
		&value.Name,
		&value.Desc,
		&value.Faulty,
	)

	if err != nil {
		return value,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return value,
		nil
}

// Выбрать значения по сервису и метрике
func (db *DB) SelectValuesByServiceAndMetric(
	service string,
	metric string,
) (
	[]Value,
	error,
) {

	errMsg := "can't select values by service and metric: "

	query := `
		SELECT *
		FROM value
		WHERE service = ?
		AND metric = ?;
	`

	rows, err := db.Query(
		query,
		service,
		metric,
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

	var values []Value

	for rows.Next() {

		var value Value

		err = rows.Scan(
			&value.Service,
			&value.Metric,
			&value.Name,
			&value.Desc,
			&value.Faulty,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		values = append(
			values,
			value,
		)
	}

	return values,
		nil
}

// Обновить значение в таблице значений
func (db *DB) UpdateValue(
	v Value,
) error {

	errMsg := "can't update value: "

	query := `
		UPDATE value
		SET desc = ?,
		faulty = ?
		WHERE service = ?
		AND metric = ?
		AND name = ?;
	`

	_, err := db.Exec(
		query,
		v.Desc,
		v.Faulty,
		v.Service,
		v.Metric,
		v.Name,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Удалить значение из таблицы значений
func (db *DB) DeleteValue(
	service string,
	metric string,
	name string,
) error {

	errMsg := "can't delete value: "

	query := `
		DELETE FROM value
		WHERE service = ?
		AND metric = ?
		AND name = ?;
	`

	_, err := db.Exec(
		query,
		service,
		metric,
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
