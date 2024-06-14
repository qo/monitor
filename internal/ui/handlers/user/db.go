package user

import "github.com/qo/monitor/internal/db"

var _db *db.DB
var err error

func init() {
	_db, _, err = db.Open()
}

func users() ([]db.User, error) {

	if err != nil {
		return nil, err
	}

	return _db.SelectAllUsers()
}

func user(username string) (db.User, error) {

	if err != nil {
		return db.User{}, err
	}

	return _db.SelectUserByUsername(
		username,
	)
}

func insert(user db.User) error {

	if err != nil {
		return err
	}

	return _db.InsertUser(user)
}

func remove(username string) error {

	if err != nil {
		return err
	}

	return _db.DeleteUser(
		username,
	)
}
