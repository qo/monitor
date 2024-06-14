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

	// Пересоздать БД
	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Открыть БД
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

	// Создать таблицы
	err = _db.CreateTables()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Валидный мессенджер
	validMessenger := db.Messenger{
		Name: "telegram",
	}

	// Невалидный мессенджер
	invalidMessenger := db.Messenger{
		Name: "...",
	}

	// Добавить валидный мессенджер
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

	// Валидный пользователь
	validUser := db.User{
		Username: "admin",
	}

	// Невалидный пользователь
	invalidUser := db.User{
		Username: "...",
	}

	// Добавить валидного пользователя
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

	// Валидный эндпоинт
	validEndpoint := db.Endpoint{
		Messenger: validMessenger.Name,
		User:      validUser.Username,
		Id:        "1",
	}

	// Невалидные эндпоинты
	invalidEndpoints := []db.Endpoint{

		// Дупликат валидного эндпоинта.
		// Является невалидным, так как
		// не выполняется ограничение PRIMARY KEY
		validEndpoint,

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

		// Является невалидным, так как
		// не выполняется условие
		// на непустую строку в Id
		{
			Messenger: validMessenger.Name,
			User:      validUser.Username,
			Id:        "",
		},
	}

	// Добавить валидный эндпоинт
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

	// Добавить невалидные эндпоинты
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

	// Пересоздать БД
	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Открыть БД
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

	// Создать таблицы
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

	// Добавить мессенджер
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

	// Добавить пользователя
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

	// Добавить эндпоинт
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

	// Выбрать все эндпоинты
	endpoints, err := _db.SelectAllEndpoints()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Если длина списка эндпоинтов не равна 1,
	// завершить работу с ошибкой
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

	// Единственный выбранный эндпоинт
	selected := endpoints[0]

	// Если эндпоинты не равны,
	// завершить работу с ошибкой
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

	// Пересоздать БД
	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Открыть БД
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

	// Создать таблицы
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

	// Добавить мессенджер
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

	// Добавить пользователя
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

	// Добавить эндпоинт
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

	// Выбрать эндпоинт
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

	// Если выбранный и добавленный эндпоинт
	// не совпадают, завершить работу с ошибкой
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

	// Пересоздать БД
	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Открыть БД
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

	// Создать таблицы
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

	// Добавить мессенджер
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

	// Добавить пользователя
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

	// Добавить эндпоинт
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

	// Выбрать эндпоинты по мессенджеру
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

	// Если длина списка выбранных эндпоинтов не равна 1,
	// завершить работу с ошибкой
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

	// Единственный выбранный эндпоинт
	selected := endpoints[0]

	// Если выбранный и добавленный
	// эндпоинты не равны,
	// завершить работу с ошибкой
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

	// Пересоздать БД
	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Открыть БД
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

	// Создать таблицы
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

	// Добавить мессенджер
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

	// Добавить пользователя
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

	// Добавить эндпоинт
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

	// Указать в качестве идентификатора эндпоинта новое значение
	endpoint.Id = "2"

	// Обновить эндпоинт
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

	// Выбрать эндпоинт
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

	// Если выбранный и добавленный
	// эндпоинты не равны,
	// завершить работу с ошибкой
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

	// Пересоздать БД
	err := setupDb()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Открыть БД
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

	// Создать таблицы
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

	// Добавить мессенджер
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

	// Добавить пользователя
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

	// Добавить эндпоинт
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

	// Удалить эндпоинт
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

	// Выбрать все эндпоинты
	endpoints, err := _db.SelectAllEndpoints()
	if err != nil {
		t.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Если длина списка
	// выбранных эндпоинтов не равна 0,
	// завершить работу с ошибкой
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
