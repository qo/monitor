package endpoint

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

// Функция, которая запустится
// при импорте пакета
func init() {
	_db, _, err = db.Open()
}

func endpoints() (
	[]db.Endpoint,
	error,
) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllEndpoints()
}

func endpoint(
	messenger string,
	user string,
) (
	db.Endpoint,
	error,
) {

	if err != nil {
		return db.Endpoint{}, err
	}

	return _db.SelectEndpointByMessengerAndUser(
		messenger,
		user,
	)
}

func update(
	endpoint db.Endpoint,
) error {

	if err != nil {
		return err
	}

	return _db.UpdateEndpoint(
		endpoint,
	)
}

func insert(
	endpoint db.Endpoint,
) error {

	if err != nil {
		return err
	}

	return _db.InsertEndpoint(
		endpoint,
	)
}

func remove(
	messenger string,
	name string,
) error {

	if err != nil {
		return err
	}

	return _db.DeleteEndpoint(
		messenger,
		name,
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

func users() (
	[]db.User,
	error,
) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllUsers()
}
