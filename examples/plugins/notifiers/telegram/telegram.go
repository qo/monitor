package telegram

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/qo/monitor/examples/plugins/notifiers/telegram/config"
	"github.com/qo/monitor/internal/db"
	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

const messengerName = "telegram"

// Данный плагин предназначен для того, чтобы
// через мессенджер Telegram
// уведомлять пользователей о статусах сервисов,
// и давать им возможность брать задачу
// в работу и помечать её как завершенную

// TODO:
// Вообще говоря, следует описать интерфейс,
// реализующий только необходимые для плагина методы,
// но можно и просто импортировать пакет db.

// Тип, который должен реализовать интерфейс
// plugins.Notifier
type notifier struct{}

// Метод Messenger() возвращает
// имя мессенджера, которое указано в БД.
// Это позволяет привязать данный плагин
// к мессенджеру из БД без хранения кода
// плагина непосредственно в БД.
func (n notifier) Messenger() string {
	return messengerName
}

// Вспомогательный тип, который
// будет использоваться, чтобы
// был доступен метод Recipient,
// возвращающий идентификатор
// эндпоинта
type recipient struct {
	endpoint db.Endpoint
}

func (r recipient) Recipient() string {
	return r.endpoint.Id
}

var (

	// Экземпляр сущности Telegram-бот
	b = &tele.Bot{}

	// Разметка для взятия задачи в работу
	takeMarkup = &tele.ReplyMarkup{}
	btnTake    = takeMarkup.Data(
		"Take",
		"take",
	)

	// Разметка для завершения задачи
	completeMarkup = &tele.ReplyMarkup{}
	btnComplete    = completeMarkup.Data(
		"Complete",
		"complete",
	)

	// Экземпляр базы данных
	_db *db.DB
)

// Запускается при первом импортировании пакета
func init() {

	errMsg := "can't init telegram plugin package: "

	// Считать конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}

	// Задать настройки Telegram-бота
	pref := telebot.Settings{
		Token: cfg.Token,
		Poller: &tele.LongPoller{
			Timeout: cfg.Timeout,
		},
	}

	// Добавить кнопки в разметку
	takeMarkup.Inline(
		takeMarkup.Row(btnTake),
	)
	completeMarkup.Inline(
		completeMarkup.Row(btnComplete),
	)

	// Создать Telegram-бота
	b, err = tele.NewBot(pref)
	if err != nil {
		log.Fatal(
			errors.New(
				errMsg +
					"can't create new telegram bot: " +
					err.Error(),
			),
		)
	}

	// Открыть базу данных
	_db, _, err = db.Open()
	if err != nil {
		log.Fatal(
			errors.New(
				errMsg +
					err.Error(),
			),
		)
	}
}

// Составить сообщение
func composeMsg(
	value db.Value,
) string {
	return fmt.Sprintf(
		`Service: %s
Metric: %s
Value: %s`,
		value.Service,
		value.Metric,
		value.Name,
	)
}

// Разобрать сообщение
func parseMsg(
	msg string,
) (
	string,
	string,
	string,
	error,
) {

	errMsg := "can't parse message: "

	var service string
	var metric string
	var value string
	_, err := fmt.Sscanf(
		msg,
		`Service: %s
Metric: %s
Value: %s`,
		&service,
		&metric,
		&value,
	)
	if err != nil {
		return service,
			metric,
			value,
			errors.New(
				errMsg +
					err.Error(),
			)
	}

	return service,
		metric,
		value,
		nil
}

// Отослать сообщение о значении
// на указанный эндпоинт
func (n notifier) Notify(
	value db.Value,
	endpoint db.Endpoint,
) error {

	r := recipient{
		endpoint,
	}

	// Составить сообщение
	msg := composeMsg(
		value,
	)

	// Если значение является ошибочным,
	// прикрепить разметку с возможностью
	// взять задачу в работу; в обратном
	// случае просто отослать сообщение
	if value.Faulty {
		b.Send(
			r,
			msg,
			takeMarkup,
		)
	} else {
		b.Send(
			r,
			msg,
		)
	}

	return nil
}

// Обработчик, который должен быть
// зарегистрирован на кнопку
// для взятия задачи в работу
func handleAssignTask(
	c tele.Context,
) error {

	errMsg := "can't handle assign task: "

	// Получить ID пользователя, который нажал на кнопку
	id := c.Sender().ID

	// Преобразовать ID в строку
	idStr := strconv.Itoa(int(id))

	// Получить эндпоинт по указанному ID
	endpoint, err := _db.SelectEndpointByMessengerAndId(
		messengerName,
		idStr,
	)
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Получить текстовое содержание сообщения
	msg := c.Message().Text

	// Разобрать сообщение
	service, metric, value, err := parseMsg(msg)
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Получить задачу для данных сервиса и метрики
	task, err := _db.SelectTaskByServiceAndMetric(
		service,
		metric,
	)
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Указать в качестве работника над задачей
	// имя пользователя, нажавшего на кнопку
	task.Worker = endpoint.User

	// Обновить задачу
	err = _db.UpdateTask(task)
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Триггеры, относящиеся к данному значению
	relatedTriggers, err := _db.SelectTriggersByServiceMetricAndValue(
		service,
		metric,
		value,
	)
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Пользователи, заинтересованные
	// в оповещении о данном значении
	var interestedUsers []string
	for _, trigger := range relatedTriggers {
		interestedUsers = append(
			interestedUsers,
			trigger.User,
		)
	}

	// Эндпоинты пользователей,
	// заинтересованных в данном значении
	var relatedEndpoints []db.Endpoint
	for _, user := range interestedUsers {

		endpoint, err := _db.SelectEndpointByMessengerAndUser(
			messengerName,
			user,
		)
		if err != nil {
			return errors.New(
				errMsg +
					err.Error(),
			)
		}

		relatedEndpoints = append(
			relatedEndpoints,
			endpoint,
		)
	}

	for _, endpoint := range relatedEndpoints {

		_, err = b.Send(
			recipient{
				endpoint,
			},
			fmt.Sprintf(
				`%s
Worker: %s`,
				c.Message().Text,
				task.Worker,
			),
			completeMarkup,
		)
		if err != nil {
			return errors.New(
				errMsg +
					"can't send message after take: " +
					err.Error(),
			)
		}
	}

	return nil
}

// Обработчик, который должен быть
// зарегистрирован на кнопку
// для завершения задачи
func handleCompleteTask(
	c tele.Context,
) error {

	errMsg := "can't handle complete task: "

	// Получить текстовое содержимое сообщения
	msg := c.Message().Text

	// Разобрать сообщение на сервис и метрику
	service, metric, value, err := parseMsg(msg)
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Удалить задачу, связанную с данными
	// сервисом и метрикой
	err = _db.DeleteTask(
		service,
		metric,
	)
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Триггеры, относящиеся к данному значению
	relatedTriggers, err := _db.SelectTriggersByServiceMetricAndValue(
		service,
		metric,
		value,
	)
	if err != nil {
		return errors.New(
			errMsg +
				err.Error(),
		)
	}

	// Пользователи, заинтересованные
	// в оповещении о данном значении
	var interestedUsers []string
	for _, trigger := range relatedTriggers {
		interestedUsers = append(
			interestedUsers,
			trigger.User,
		)
	}

	// Эндпоинты пользователей,
	// заинтересованных в данном значении
	var relatedEndpoints []db.Endpoint
	for _, user := range interestedUsers {

		endpoint, err := _db.SelectEndpointByMessengerAndUser(
			messengerName,
			user,
		)
		if err != nil {
			return errors.New(
				errMsg +
					err.Error(),
			)
		}

		relatedEndpoints = append(
			relatedEndpoints,
			endpoint,
		)
	}

	for _, endpoint := range relatedEndpoints {

		_, err = b.Send(
			recipient{
				endpoint,
			},
			fmt.Sprintf(
				`%s
Status: Completed`,
				c.Message().Text,
			),
		)
		if err != nil {
			return errors.New(
				errMsg +
					"can't send message after complete: " +
					err.Error(),
			)
		}
	}

	return nil
}

// Метод, запускающий обработчики для
// взятия задачи в работу
// и завершения задачи
func (n notifier) Run() {
	b.Handle(
		&btnTake,
		handleAssignTask,
	)
	b.Handle(
		&btnComplete,
		handleCompleteTask,
	)
	go b.Start()
}

// Экспортировать переменную типа notifier.
// Это можно назвать паттерном синглтон,
// так как тип notifier не экспортирован,
// и как следствие переменные данного
// типа не могут быть созданы напрямую
// вне данного пакета.
var Notifier notifier
