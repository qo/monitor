package db_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/qo/monitor/internal/db"
)

// Это пример того, как провести тестирование
// функций для работы с определенной таблицей.

// Так как написание тестов для всех таблиц
// является очень времязатратным занятием,
// на данный момент оно опущено.

// TODO: написать тесты для всех таблиц.

// Являются ли эндпоинты одинаковыми
func equal(
	first db.Endpoint,
	second db.Endpoint,
) bool {
	if first.Messenger != second.Messenger {
		return false
	}
	if first.Id != second.Id {
		return false
	}
	if first.User != second.User {
		return false
	}
	if first.Desc != second.Desc {
		return false
	}
	return true
}

// Протестировать функцию
// добавления эндпоинта в таблицу эндпоинтов
func TestInsertEndpoint(
	t *testing.T,
) {

	errMsg := "can't test insert endpoint: "

	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	_db, closer, err := db.Open()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}
	defer closer()

	err = _db.CreateTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	validMessenger := db.Messenger{
		Name: "telegram",
	}

	invalidMessenger := db.Messenger{
		Name: "...",
	}

	err = _db.InsertMessenger(
		validMessenger,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	validUser := db.User{
		Username: "admin",
	}

	invalidUser := db.User{
		Username: "...",
	}

	err = _db.InsertUser(
		validUser,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	validEndpoint := db.Endpoint{
		Messenger: validMessenger.Name,
		User:      validUser.Username,
		Id:        "1",
	}

	invalidEndpoints := []db.Endpoint{
		validEndpoint, // duplicate
		{
			Messenger: invalidMessenger.Name,
			User:      invalidUser.Username,
			Id:        "1",
		},
		{
			Messenger: validMessenger.Name,
			User:      invalidUser.Username,
			Id:        "1",
		},
		{
			Messenger: invalidMessenger.Name,
			User:      validUser.Username,
			Id:        "1",
		},
		{
			Messenger: validMessenger.Name,
			User:      validUser.Username,
			Id:        "",
		},
	}

	err = _db.InsertEndpoint(
		validEndpoint,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	for _, invalidEndpoint := range invalidEndpoints {
		err = _db.InsertEndpoint(
			invalidEndpoint,
		)
		if err == nil {
			t.Fatal(
				errors.New(
					errMsg +
						"invalid endpoint passed: " +
						fmt.Sprintf(
							"%+v",
							invalidEndpoint,
						),
				),
			)
		}
	}
}

// Протестировать функцию
// выбора всех эндпоинтов
func TestSelectAllEndpoints(
	t *testing.T,
) {

	errMsg := "can't test select all endpoints: "

	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	_db, closer, err := db.Open()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}
	defer closer()

	err = _db.CreateTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	messenger := db.Messenger{
		Name: "telegram",
	}

	err = _db.InsertMessenger(
		messenger,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	user := db.User{
		Username: "admin",
	}

	err = _db.InsertUser(
		user,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	endpoint := db.Endpoint{
		Messenger: messenger.Name,
		User:      user.Username,
		Id:        "1",
	}

	err = _db.InsertEndpoint(
		endpoint,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	endpoints, err := _db.SelectAllEndpoints()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	if len(endpoints) != 1 {
		t.Fatal(
			errors.New(
				errMsg +
					fmt.Sprintf(
						"length of selected endpoints is %d instead of %d",
						len(endpoints),
						1,
					),
			),
		)
	}

	selected := endpoints[0]
	if !equal(
		endpoint,
		selected,
	) {
		t.Fatal(
			errors.New(
				errMsg +
					fmt.Sprintf(
						"expected %+v but got %+v",
						endpoint,
						selected,
					),
			),
		)
	}
}

// Протестировать функцию
// выбора эндпоинта по мессенджеру и пользователю
func TestSelectEndpointByMessengerAndUser(
	t *testing.T,
) {

	errMsg := "can't test select endpoint by messenger and user: "

	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	_db, closer, err := db.Open()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}
	defer closer()

	err = _db.CreateTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	messenger := db.Messenger{
		Name: "telegram",
	}

	err = _db.InsertMessenger(
		messenger,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	user := db.User{
		Username: "admin",
	}

	err = _db.InsertUser(
		user,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	endpoint := db.Endpoint{
		Messenger: messenger.Name,
		User:      user.Username,
		Id:        "1",
	}

	err = _db.InsertEndpoint(
		endpoint,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	selected, err := _db.SelectEndpointByMessengerAndUser(
		messenger.Name,
		user.Username,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	if !equal(
		endpoint,
		selected,
	) {
		t.Fatal(
			errors.New(
				errMsg +
					fmt.Sprintf(
						"expected %+v but got %+v",
						endpoint,
						selected,
					),
			),
		)
	}
}

// Протестировать функцию
// выбора эндпоинтов по мессенджерам
func TestSelectEndpointsByMessengers(
	t *testing.T,
) {

	errMsg := "can't test select endpoint by messenger and user: "

	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	_db, closer, err := db.Open()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}
	defer closer()

	err = _db.CreateTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	messenger := db.Messenger{
		Name: "telegram",
	}

	err = _db.InsertMessenger(
		messenger,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	user := db.User{
		Username: "admin",
	}

	err = _db.InsertUser(
		user,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	endpoint := db.Endpoint{
		Messenger: messenger.Name,
		User:      user.Username,
		Id:        "1",
	}

	err = _db.InsertEndpoint(
		endpoint,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	endpoints, err := _db.SelectEndpointsByMessenger(
		messenger.Name,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	if len(endpoints) != 1 {
		t.Fatal(
			errors.New(
				errMsg +
					fmt.Sprintf(
						"length of selected endpoints is %d instead of %d",
						len(endpoints),
						1,
					),
			),
		)
	}

	selected := endpoints[0]

	if !equal(
		endpoint,
		selected,
	) {
		t.Fatal(
			errors.New(
				errMsg +
					fmt.Sprintf(
						"expected %+v but got %+v",
						endpoint,
						selected,
					),
			),
		)
	}
}

// Протестировать функцию
// обновления эндпоинта в таблице эндпоинтов
func TestUpdateEndpoint(
	t *testing.T,
) {

	errMsg := "can't test select endpoint by messenger and user: "

	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	_db, closer, err := db.Open()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}
	defer closer()

	err = _db.CreateTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	messenger := db.Messenger{
		Name: "telegram",
	}

	err = _db.InsertMessenger(
		messenger,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	user := db.User{
		Username: "admin",
	}

	err = _db.InsertUser(
		user,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	endpoint := db.Endpoint{
		Messenger: messenger.Name,
		User:      user.Username,
		Id:        "1",
	}

	err = _db.InsertEndpoint(
		endpoint,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	endpoint.Id = "2"

	err = _db.UpdateEndpoint(
		endpoint,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	selected, err := _db.SelectEndpointByMessengerAndUser(
		messenger.Name,
		user.Username,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	if !equal(
		endpoint,
		selected,
	) {
		t.Fatal(
			errors.New(
				errMsg +
					fmt.Sprintf(
						"expected %+v but got %+v",
						endpoint,
						selected,
					),
			),
		)
	}
}

// Протестировать функцию
// удаления эндпоинта из таблицы эндпоинтов
func TestDeleteEndpoint(
	t *testing.T,
) {

	errMsg := "can't test select endpoint by messenger and user: "

	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	_db, closer, err := db.Open()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}
	defer closer()

	err = _db.CreateTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	messenger := db.Messenger{
		Name: "telegram",
	}

	err = _db.InsertMessenger(
		messenger,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	user := db.User{
		Username: "admin",
	}

	err = _db.InsertUser(
		user,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	endpoint := db.Endpoint{
		Messenger: messenger.Name,
		User:      user.Username,
		Id:        "1",
	}

	err = _db.InsertEndpoint(
		endpoint,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	err = _db.DeleteEndpoint(
		messenger.Name,
		user.Username,
	)
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	endpoints, err := _db.SelectAllEndpoints()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	if len(endpoints) > 0 {
		t.Fatal(
			errors.New(
				errMsg +
					fmt.Sprintf(
						"length of endpoints is equal to %d instead of %d",
						len(endpoints),
						0,
					),
			),
		)
	}
}
