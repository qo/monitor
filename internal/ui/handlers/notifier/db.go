package notifier

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

// Функция, которая запустится
// при импорте пакета
func init() {
	_db, _, err = db.Open()
}

func notifiers() (
	[]db.Notifier,
	error,
) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllNotifiers()
}

func notifier(
	messenger string,
) (
	db.Notifier,
	error,
) {

	if err != nil {
		return db.Notifier{}, err
	}

	return _db.SelectNotifierByMessenger(
		messenger,
	)
}

func insert(
	notifier db.Notifier,
) error {

	if err != nil {
		return err
	}

	return _db.InsertNotifier(
		notifier,
	)
}

func remove(
	messenger string,
) error {

	if err != nil {
		return err
	}

	return _db.DeleteNotifier(
		messenger,
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
