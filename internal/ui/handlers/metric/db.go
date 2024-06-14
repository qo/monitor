package metric

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

func init() {
	_db, _, err = db.Open()
}

func metrics() ([]db.Metric, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllMetrics()
}

func metric(service, name string) (db.Metric, error) {

	if err != nil {
		return db.Metric{}, err
	}

	return _db.SelectMetricByServiceAndName(
		service,
		name,
	)
}

func update(metric db.Metric) error {

	if err != nil {
		return err
	}

	return _db.UpdateMetric(metric)
}

func insert(metric db.Metric) error {

	if err != nil {
		return err
	}

	return _db.InsertMetric(metric)
}

func remove(service, name string) error {

	if err != nil {
		return err
	}

	return _db.DeleteMetric(
		service,
		name,
	)
}

func services() ([]db.Service, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllServices()
}
