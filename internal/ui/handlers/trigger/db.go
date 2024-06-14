package trigger

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

func init() {
	_db, _, err = db.Open()
}

func triggers() ([]db.Trigger, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllTriggers()
}

func trigger(
	service string,
	metric string,
	value string,
	messenger string,
	user string,
) (
	db.Trigger,
	error,
) {

	if err != nil {
		return db.Trigger{}, err
	}

	return _db.SelectTriggerByServiceMetricValueMessengerAndUser(
		service,
		metric,
		value,
		messenger,
		user,
	)
}

// func update(trigger db.Trigger) error {

// 	if err != nil {
// 		return err
// 	}

// 	return _db.UpdateTrigger(trigger)
// }

func insert(
	trigger db.Trigger,
) error {

	if err != nil {
		return err
	}

	return _db.InsertTrigger(
		trigger,
	)
}

func remove(
	service string,
	metric string,
	value string,
	messenger string,
	user string,
) error {

	if err != nil {
		return err
	}

	return _db.DeleteTrigger(
		service,
		metric,
		value,
		messenger,
		user,
	)
}

func services() (
	[]db.Service,
	error,
) {

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

func values(
	service,
	metric string,
) (
	[]db.Value,
	error,
) {

	if err != nil {
		return nil, err
	}

	return _db.SelectValuesByServiceAndMetric(
		service,
		metric,
	)
}

func messengers() (
	[]db.Messenger,
	error,
) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllMessengers()
}

func endpoints(
	messenger string,
) (
	[]db.Endpoint,
	error,
) {

	if err != nil {
		return nil, err
	}

	return _db.SelectEndpointsByMessenger(
		messenger,
	)
}
