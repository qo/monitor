package service

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

func init() {
	_db, _, err = db.Open()
}

func services() ([]db.Service, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllServices()
}

func service(name string) (db.Service, error) {

	if err != nil {
		return db.Service{}, err
	}

	return _db.SelectServiceByName(
		name,
	)
}

func update(service db.Service) error {

	if err != nil {
		return err
	}

	return _db.UpdateService(service)
}

func insert(service db.Service) error {

	if err != nil {
		return err
	}

	return _db.InsertService(service)
}

func remove(name string) error {

	if err != nil {
		return err
	}

	return _db.DeleteService(
		name,
	)
}

