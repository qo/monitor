package poller

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

func init() {
	_db, _, err = db.Open()
}

func pollers() ([]db.Poller, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllPollers()
}

func poller(
	service string,
	metric string,
) (
	db.Poller,
	error,
) {

	if err != nil {
		return db.Poller{}, err
	}

	return _db.SelectPollerByServiceAndMetric(
		service,
		metric,
	)
}

func insert(
	poller db.Poller,
) error {

	if err != nil {
		return err
	}

	return _db.InsertPoller(
		poller,
	)
}

func remove(
	service string,
	metric string,
) error {

	if err != nil {
		return err
	}

	return _db.DeletePoller(
		service,
		metric,
	)
}

func services() ([]db.Service, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllServices()
}

func metrics(
	service string,
) (
	[]db.Metric,
	error,
) {

	if err != nil {
		return nil, err
	}

	return _db.SelectMetricsByService(
		service,
	)
}
