package db

import (
	"errors"
	"strconv"
)

type Task struct {
	Service string
	Metric  string
	Value   string
	Worker  string
}

// Создать таблицу задач
func (db *DB) CreateTasks() error {

	errMsg := "can't create task table: "

	query := `
		CREATE TABLE task(
			service VARCHAR(32) NOT NULL,
			metric VARCHAR(32) NOT NULL,
			value VARCHAR(32) NOT NULL,
			worker VARCHAR(32),
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
			FOREIGN KEY (worker)
			REFERENCES user(username),
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

// Удалить таблицу задач
func (db *DB) DropTasks() error {

	errMsg := "can't drop task table: "

	query := `
		DROP TABLE task;
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

// Добавить задачу в таблицу задач
func (db *DB) InsertTask(
	t Task,
) error {

	errMsg := "can't insert task: "

	query := `
		INSERT INTO task
		VALUES(
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
		t.Worker,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Добавить задачу без работника в таблицу задач
func (db *DB) InsertTaskWithNoWorker(
	t Task,
) error {

	errMsg := "can't insert task with no worker: "

	query := `
		INSERT INTO task
		VALUES(
			?,
			?,
			?,
			NULL
		);
	`

	_, err := db.Exec(
		query,
		t.Service,
		t.Metric,
		t.Value,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Обновить задачу в таблице задач
func (db *DB) UpdateTask(
	t Task,
) error {

	errMsg := "can't update task: "

	query := `
		UPDATE task
		SET value = ?,
		worker = ?
		WHERE service = ?
		AND metric = ?;
	`

	_, err := db.Exec(
		query,
		t.Value,
		t.Worker,
		t.Service,
		t.Metric,
	)

	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	return nil
}

// Выбрать все задачи
func (db *DB) SelectAllTasks() (
	[]Task,
	error,
) {

	errMsg := "can't select tasks: "

	query := `
		SELECT *
		FROM task;
	`

	rows, err := db.Query(
		query,
	)

	if err != nil {
		return nil, errors.New(
			errMsg +
				"can't query: " +
				err.Error(),
		)
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {

		var task Task

		// Переменная-пустышка,
		// необходимая для сканирования
		// ненужных полей
		var dummy any

		err = rows.Scan(
			&task.Service,
			&task.Metric,
			&task.Value,
			&dummy,
		)

		if err != nil {
			return nil, errors.New(
				errMsg +
					"can't scan: " +
					err.Error(),
			)
		}

		// Отдельно просканировать атрибут Worker,
		// не проверяя ошибку,
		// так как он может быть равен NULL
		rows.Scan(
			&dummy,
			&dummy,
			&dummy,
			&task.Worker,
		)

		tasks = append(
			tasks,
			task,
		)
	}

	return tasks,
		nil
}

// Выбрать задачу по сервису и метрике
func (db *DB) SelectTaskByServiceAndMetric(
	service string,
	metric string,
) (
	Task,
	error,
) {

	errMsg := "can't select task: "

	query := `
		SELECT
			service,
			metric,
			value,
			COALESCE(worker, '')
		FROM task
		WHERE service = ?
		AND metric = ?;
	`

	row := db.QueryRow(
		query,
		service,
		metric,
	)

	var task Task

	// Переменная-пустышка,
	// необходимая для сканирования
	// ненужных полей
	var dummy any

	err := row.Scan(
		&task.Service,
		&task.Metric,
		&task.Value,
		&dummy,
	)

	if err != nil {
		return task, errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Отдельно просканировать атрибут Worker,
	// не проверяя ошибку,
	// так как он может быть равен NULL
	row.Scan(
		&dummy,
		&dummy,
		&dummy,
		&task.Worker,
	)

	return task,
		nil
}

// Удалить задачу из таблицы задач
func (db *DB) DeleteTask(
	service string,
	metric string,
) error {

	errMsg := "can't delete task: "

	query := `
		DELETE FROM task
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

// Существует ли задача
func (db *DB) TaskExists(
	service string,
	metric string,
) (
	bool,
	error,
) {

	errMsg := "can't check if task exists: "

	query := `
		SELECT COUNT(*)
		FROM task
		WHERE service = ?
		AND metric = ?;
	`

	row := db.QueryRow(
		query,
		service,
		metric,
	)

	var count int

	err := row.Scan(
		&count,
	)

	if err != nil {
		return false,
			errors.New(
				errMsg +
					"can't scan: " +
					err.Error(),
			)
	}

	if count == 0 {
		return false, nil
	} else if count == 1 {
		return true, nil
	}

	return false,
		errors.New(
			errMsg +
				"count is neither 0 nor 1 but " +
				strconv.Itoa(count),
		)
}
