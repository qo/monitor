package task

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

func init() {
	_db, _, err = db.Open()
}

func tasks() ([]db.Task, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllTasks()
}

func task(service, metric string) (db.Task, error) {

	if err != nil {
		return db.Task{}, err
	}

	return _db.SelectTaskByServiceAndMetric(
		service,
		metric,
	)
}

func update(task db.Task) error {

	if err != nil {
		return err
	}

	return _db.UpdateTask(task)
}

func insert(task db.Task) error {

	if err != nil {
		return err
	}

	return _db.InsertTask(task)
}

func remove(service, metric string) error {

	if err != nil {
		return err
	}

	return _db.DeleteTask(
		service,
		metric,
	)
}

func users() ([]db.User, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllUsers()
}
