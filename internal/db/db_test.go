package db_test

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/qo/monitor/internal/db"
)

var dbName = "monitor.db"

var selectTablesQuery = `
	SELECT name
	FROM sqlite_master
	WHERE type = 'table';
`

// Протестировать функцию открытия БД
func TestOpen(
	t *testing.T,
) {

	errMsg := "can't test open: "

	_db, _, err := db.Open()

	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					"can't open: " +
					err.Error(),
			),
		)
	}

	_, err = _db.Query(
		selectTablesQuery,
	)

	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					"can't query: " +
					err.Error(),
			),
		)
	}
}

// Пересоздать файл БД
func setupDb() error {

	errMsg := "can't setup db: "

	_, err := os.Stat(
		dbName,
	)

	if err == nil {

		err = os.Remove(
			dbName,
		)
		if err != nil {
			return errors.New(
				errMsg +
					"can't remove: " +
					err.Error(),
			)
		}

		_, err = os.Create(
			dbName,
		)
		if err != nil {
			return errors.New(
				errMsg +
					"can't create: " +
					err.Error(),
			)
		}

	} else if errors.Is(err, os.ErrNotExist) {

		_, err = os.Create(
			dbName,
		)
		if err != nil {
			return errors.New(
				errMsg +
					"can't create: " +
					err.Error(),
			)
		}

	} else {
		return errors.New(
			errMsg +
				"can't stat: " +
				err.Error(),
		)
	}

	return nil
}

// Являются ли два среза строк равными
// https://stackoverflow.com/a/36000696
func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		diff[_x]++
	}
	for _, _y := range y {
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}

// Протестировать функцию создания всех таблиц
func TestCreateTables(
	t *testing.T,
) {

	errMsg := "can't test create: "

	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	_db, _, err := db.Open()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// err = _db.CreateTables()
	// if err != nil {
	// 	t.Fatal(
	// 		errors.New(
	// 			errMsg +
	// 				err.Error(),
	// 		),
	// 	)
	// }

	// err = _db.DropTables()
	// if err != nil {
	// 	t.Fatal(
	// 		errors.New(
	// 			errMsg +
	// 				err.Error(),
	// 		),
	// 	)
	// }

	err = _db.CreateTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	rows, err := _db.Query(
		selectTablesQuery,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					"can't query: " +
					err.Error(),
			),
		)
	}

	var tables []string

	for rows.Next() {

		var table string

		err = rows.Scan(
			&table,
		)

		if err != nil {
			t.Fatal(
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				),
			)
		}

		tables = append(
			tables,
			table,
		)
	}

	want := []string{
		"service",
		"metric",
		"value",
		"messenger",
		"endpoint",
		"trigger",
		"user",
		"task",
		"poller",
		"notifier",
	}

	if !sameStringSlice(tables, want) {
		t.Fatal(
			errors.New(
				errMsg +
					"wanted and actual tables don't match: (want: " +
					strings.Join(
						want,
						",",
					) +
					"actual: " +
					strings.Join(
						tables,
						",",
					) +
					")",
			),
		)
	}
}

// Протестировать функцию удаления всех таблиц
func TestDropTables(
	t *testing.T,
) {

	errMsg := "can't test drop: "

	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	_db, _, err := db.Open()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					"can't open: " +
					err.Error(),
			),
		)
	}

	err = _db.CreateTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					"can't create: " +
					err.Error(),
			),
		)
	}

	err = _db.DropTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					"can't drop: " +
					err.Error(),
			),
		)
	}

	rows, err := _db.Query(
		selectTablesQuery,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					"can't query: " +
					err.Error(),
			),
		)
	}

	var tables []string

	for rows.Next() {

		var table string

		err = rows.Scan(
			&table,
		)

		if err != nil {
			t.Fatal(
				errors.New(
					errMsg +
						"can't scan: " +
						err.Error(),
				),
			)
		}

		tables = append(
			tables,
			table,
		)
	}

	if len(tables) > 0 {
		t.Fatal(
			errors.New(
				errMsg +
					"there are " +
					strconv.Itoa(
						len(tables),
					) +
					"tables after dropping: " +
					strings.Join(
						tables,
						",",
					),
			),
		)
	}
}
