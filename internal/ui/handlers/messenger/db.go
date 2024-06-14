package messenger

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

func init() {
	_db, _, err = db.Open()
}

func messengers() ([]db.Messenger, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllMessengers()
}

func messenger(name string) (db.Messenger, error) {

	if err != nil {
		return db.Messenger{}, err
	}

	return _db.SelectMessengerByName(
		name,
	)
}

func update(messenger db.Messenger) error {

	if err != nil {
		return err
	}

	return _db.UpdateMessenger(messenger)
}

func insert(messenger db.Messenger) error {

	if err != nil {
		return err
	}

	return _db.InsertMessenger(messenger)
}

func remove(name string) error {

	if err != nil {
		return err
	}

	return _db.DeleteMessenger(
		name,
	)
}
