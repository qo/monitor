package db

import "errors"

type Metric struct {
	Service string
	Name    string
	Desc    string
}

// Создать таблицу метрик
func (db *DB) CreateMetrics() error {

	errMsg := "can't create metric table: "

	query := `
		CREATE TABLE metric(
			service VARCHAR(32) NOT NULL,
			name VARCHAR(32) NOT NULL,
			desc VARCHAR(256),
			FOREIGN KEY (service)
			REFERENCES service(name),
			PRIMARY KEY (service, name)
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

// Удалить таблицу метрик
func (db *DB) DropMetrics() error {

	errMsg := "can't drop metric table: "

	query := `
		DROP TABLE metric;
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

// Добавить метрику в таблицу метрик
func (db *DB) InsertMetric(
	m Metric,
) error {

	errMsg := "can't insert metric: "

	query := `
		INSERT INTO metric
		VALUES(
			?,
			?,
			?
		);
	`

	_, err := db.Exec(
		query,
		m.Service,
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

// Выбрать все метрики
func (db *DB) SelectAllMetrics() (
	[]Metric,
	error,
) {

	errMsg := "can't select all metrics: "

	query := `
		SELECT *
		FROM metric;
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

	var metrics []Metric

	for rows.Next() {

		var metric Metric

		err = rows.Scan(
			&metric.Service,
			&metric.Name,
			&metric.Desc,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		metrics = append(
			metrics,
			metric,
		)
	}

	return metrics,
		nil
}

// Выбрать метрику по сервису и имени
func (db *DB) SelectMetricByServiceAndName(
	service string,
	name string,
) (
	Metric,
	error,
) {

	errMsg := "can't select metric: "

	query := `
		SELECT *
		FROM metric
		WHERE service = ?
		AND name = ?;
	`

	row := db.QueryRow(
		query,
		service,
		name,
	)

	var metric Metric

	err := row.Scan(
		&metric.Service,
		&metric.Name,
		&metric.Desc,
	)

	if err != nil {
		return metric,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return metric,
		nil
}

// Выбрать метрики по сервису
func (db *DB) SelectMetricsByService(
	service string,
) (
	[]Metric,
	error,
) {

	errMsg := "can't select metrics by service: "

	query := `
		SELECT *
		FROM metric
		WHERE service = ?;
	`

	rows, err := db.Query(
		query,
		service,
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

	var metrics []Metric

	for rows.Next() {

		var metric Metric

		err = rows.Scan(
			&metric.Service,
			&metric.Name,
			&metric.Desc,
		)

		if err != nil {
			return nil,
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				)
		}

		metrics = append(
			metrics,
			metric,
		)
	}

	return metrics,
		nil
}

// Обновить метрику в таблице метрик
func (db *DB) UpdateMetric(
	m Metric,
) error {

	errMsg := "can't update metric: "

	query := `
		UPDATE metric
		SET desc = ?
		WHERE service = ?
		AND name = ?;
	`

	_, err := db.Exec(
		query,
		m.Desc,
		m.Service,
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

// Удалить метрику из таблицы метрик
func (db *DB) DeleteMetric(
	service string,
	name string,
) error {

	errMsg := "can't delete metric: "

	query := `
		DELETE FROM metric
		WHERE service = ?
		AND name = ?;
	`

	_, err := db.Exec(
		query,
		service,
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
