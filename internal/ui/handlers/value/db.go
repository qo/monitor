package value

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

func init() {
	_db, _, err = db.Open()
}

func values() ([]db.Value, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllValues()
}

func value(service, metric, name string) (db.Value, error) {

	if err != nil {
		return db.Value{}, err
	}

	return _db.SelectValueByServiceMetricAndName(
		service,
		metric,
		name,
	)
}

func update(value db.Value) error {

	if err != nil {
		return err
	}

	return _db.UpdateValue(value)
}

func insert(value db.Value) error {

	if err != nil {
		return err
	}

	return _db.InsertValue(value)
}

func remove(service, metric, name string) error {

	if err != nil {
		return err
	}

	return _db.DeleteValue(
		service,
		metric,
		name,
	)
}

func services() ([]db.Service, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllServices()
}

func metrics(service string) ([]db.Metric, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectMetricsByService(service)
}
