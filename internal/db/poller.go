package db

import "errors"

type Poller struct {
	Service string
	Metric  string
}

// Создать таблицу подписывающих плагинов
func (db *DB) CreatePollers() error {

	errMsg := "can't create poller table: "

	query := `
		CREATE TABLE poller(
			service VARCHAR(32) NOT NULL,
			metric VARCHAR(32) NOT NULL,
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
				metric
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

// Удалить таблицу подписывающих плагинов
func (db *DB) DropPollers() error {

	errMsg := "can't drop poller table: "

	query := `
		DROP TABLE poller;
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

// Добавить подписывающий плагин в таблицу публикующих плагинов
func (db *DB) InsertPoller(
	p Poller,
) error {

	errMsg := "can't insert poller: "

	query := `
		INSERT INTO poller
		VALUES(
			?,
			?
		);
	`

	_, err := db.Exec(
		query,
		p.Service,
		p.Metric,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать все подписывающие плагины
func (db *DB) SelectAllPollers() (
	[]Poller,
	error,
) {

	errMsg := "can't select all pollers: "

	query := `
		SELECT *
		FROM poller;
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

	var pollers []Poller

	for rows.Next() {

		var poller Poller

		err = rows.Scan(
			&poller.Service,
			&poller.Metric,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		pollers = append(
			pollers,
			poller,
		)
	}

	return pollers,
		nil
}

// Выбрать подписывающий плагин по сервису и метрике
func (db *DB) SelectPollerByServiceAndMetric(
	service string,
	metric string,
) (
	Poller,
	error,
) {

	errMsg := "can't select poller: "

	query := `
		SELECT *
		FROM poller
		WHERE service = ?
		AND metric = ?;
	`

	row := db.QueryRow(
		query,
		service,
		metric,
	)

	var poller Poller

	err := row.Scan(
		&poller.Service,
		&poller.Metric,
	)
	if err != nil {
		return poller,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return poller,
		nil
}

// Посчитать все подписывающие плагины
func (db *DB) CountAllPollers() (
	int,
	error,
) {

	errMsg := "can't count all pollers: "

	query := `
		SELECT COUNT(*)
		FROM poller;
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

// Удалить опрашивающий плагин
// из таблицы опрашивающих плагинов
func (db *DB) DeletePoller(
	service string,
	metric string,
) error {

	errMsg := "can't delete poller: "

	query := `
		DELETE FROM poller
		WHERE service = ?
		AND metric = ?;
	`

	_, err := db.Exec(
		query,
		service,
		metric,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}
