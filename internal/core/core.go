package core

import (
	"errors"
	"fmt"
	"time"

	"github.com/qo/monitor/internal/config"
	"github.com/qo/monitor/internal/db"
)

// Тип функции Value
type ValueFuncType func() (
	string,
	error,
)

// Тип функции Notify
type NotifyFuncType func(
	db.Value,
	db.Endpoint,
) error

// Тип, связывающий сущность
// опрашивающий плагин из БД
// и код этого плагина
type Poller struct {
	db.Poller
	ValueFunc ValueFuncType
}

// Задать функцию ValueFunc для типа Poller
func (p Poller) SetValueFunc(
	f ValueFuncType,
) {
	p.ValueFunc = f
}

// Тип, связывающий сущность
// публикующий плагин из БД
// и код этого плагина
type Notifier struct {
	db.Notifier
	NotifyFunc NotifyFuncType
}

// Задать функцию NotifyFunc для типа Notifier
func (n Notifier) SetNotifyFunc(
	f NotifyFuncType,
) {
	n.NotifyFunc = f
}

// Функция, запускающая ядро
func Run() error {

	_db, close, err := db.Open()
	if err != nil {
		return err
	}
	defer close()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Посчитать опрашивающие плагины в БД
	dbPollersCount, err := _db.CountAllPollers()
	if err != nil {
		return errors.New(
			"can't count all pollers: " +
				err.Error(),
		)
	}

	// Если количество плагинов в БД
	// не совпадает с количеством плагинов в коде,
	// завершить работу и логгировать ошибку
	if len(pollerPlugins) != dbPollersCount {
		return errors.New(
			fmt.Sprintf(
				"there are %d pollers in your db, but %d pollers in your source code",
				dbPollersCount,
				len(pollerPlugins),
			),
		)
	}

	// Аналогично опрашивающим плагинам
	dbNotifiersCount, err := _db.CountAllNotifiers()
	if err != nil {
		return errors.New(
			"can't count all notifiers: " +
				err.Error(),
		)
	}

	// Аналогично опрашивающим плагинам
	if len(notifierPlugins) != dbNotifiersCount {
		return errors.New(
			fmt.Sprintf(
				"there are %d notifiers in your db, but %d notifiers in your source code",
				dbNotifiersCount,
				len(notifierPlugins),
			),
		)
	}

	// Запустить метод Run для всех публикующих плагинов
	for _, notifier := range notifierPlugins {
		notifier.Run()
	}

	var pollers []Poller

	for _, pollerPlugin := range pollerPlugins {

		// Получить опрашивающий плагин из БД,
		// сервис и значение которого совпадают
		// с сервисом и значением плагина
		dbPoller, err := _db.SelectPollerByServiceAndMetric(
			pollerPlugin.Service(),
			pollerPlugin.Metric(),
		)
		if err != nil {
			return errors.New(
				"can't select poller by service and metric: " +
					err.Error(),
			)
		}

		var poller Poller

		// Скопировать атрибуты сущности из БД
		poller.Service = dbPoller.Service
		poller.Metric = dbPoller.Metric

		// Скопировать метод Value из плагина
		poller.ValueFunc = pollerPlugin.Value

		pollers = append(
			pollers,
			poller,
		)
	}

	var notifiers []Notifier

	for _, notifierPlugin := range notifierPlugins {

		// Получить публикующий плагин из БД,
		// мессенджер которого совпадает
		// с мессенджером плагина
		dbNotifier, err := _db.SelectNotifierByMessenger(
			notifierPlugin.Messenger(),
		)
		if err != nil {
			return errors.New(
				"can't select notifier by messenger: " +
					err.Error(),
			)
		}

		var notifier Notifier

		// Скопировать атрибуты сущности из БД
		notifier.Messenger = dbNotifier.Messenger

		// Скопировать метод Notify из плагина
		notifier.NotifyFunc = notifierPlugin.Notify

		notifiers = append(
			notifiers,
			notifier,
		)
	}

	// Бесконечный цикл с задержкой
	for {

		// Ассоциативный массив,
		// который ставит в соотвествие
		// значению срез функций
		// типа func() error.
		m := make(
			map[db.Value][](func() error),
			0,
		)

		// Выбрать все триггеры
		triggers, err := _db.SelectAllTriggers()
		if err != nil {
			return errors.New(
				"can't select all triggers: " +
					err.Error(),
			)
		}

		// Для каждого триггера
		for _, trigger := range triggers {
			// Для каждого опрашивающего плагина
			for _, poller := range pollers {
				// Если сервис и метрика
				// триггера и опрашивающего плагина совпадают
				if trigger.Service == poller.Service &&
					trigger.Metric == poller.Metric {
					// Для каждого публикующего плагина
					for _, notifier := range notifiers {
						// Если мессенджер
						// триггер и публикующего плагина совпадают
						if trigger.Messenger == notifier.Messenger {

							// Получить значение
							value, err := _db.SelectValueByServiceMetricAndName(
								trigger.Service,
								trigger.Metric,
								trigger.Value,
							)
							if err != nil {
								return errors.New(
									"can't select value: " +
										err.Error(),
								)
							}

							// Получить эндпоинт
							endpoint, err := _db.SelectEndpointByMessengerAndUser(
								trigger.Messenger,
								trigger.User,
							)
							if err != nil {
								return errors.New(
									"can't select value: " +
										err.Error(),
								)
							}

							f := func() error {
								err = notifier.NotifyFunc(
									value,
									endpoint,
								)
								return err
							}

							// Добавить функцию Notify
							// в ассоциативный массив
							m[value] = append(
								m[value],
								f,
							)
						}
					}
				}
			}
		}

		// Для каждого опрашивающего плагина
		for _, poller := range pollers {

			// Получить значение
			valueStr, err := poller.ValueFunc()
			if err != nil {
				return errors.New(
					"error while getting value: " +
						err.Error(),
				)
			}

			// Получить сущность значение
			value, err := _db.SelectValueByServiceMetricAndName(
				poller.Service,
				poller.Metric,
				valueStr,
			)
			if err != nil {
				return errors.New(
					"can't select value: " +
						err.Error(),
				)
			}

			// Узнать, существует ли задача,
			// связанная с данными сервисом и метрикой
			taskExists, err := _db.TaskExists(
				value.Service,
				value.Metric,
			)
			if err != nil {
				return err
			}

			// Если задача существует,
			// перейти на следующую итерацию
			if taskExists {
				continue
			}

			// Если значение ошибочное,
			// сформировать задачу
			if value.Faulty {
				err = _db.InsertTaskWithNoWorker(
					db.Task{
						Service: value.Service,
						Metric:  value.Metric,
						Value:   value.Name,
					},
				)
				if err != nil {
					return errors.New(
						"can't insert task: " +
							err.Error(),
					)
				}
			}

			// Получить функции, которые нужно выполнить
			// для данного значения
			funcs, ok := m[value]

			// Если такие функции есть
			if ok {
				// Для каждой функции
				for _, fun := range funcs {
					// Выполнить функцию
					err = fun()
					if err != nil {
						return errors.New(
							"can't call trigger func: " +
								err.Error(),
						)
					}
				}
			}
		}

		// Подождать перед следующей итерацией
		time.Sleep(
			cfg.Delay,
		)
	}
}

// Добавить тестовые данные
func InsertExampleData(_db *db.DB) error {

	err := _db.DropTables()
	if err != nil {
		return errors.New(
			"can't drop tables: " +
				err.Error(),
		)
	}

	err = _db.CreateTables()
	if err != nil {
		return errors.New(
			"can't create tables: " +
				err.Error(),
		)
	}

	s := db.Service{
		Name: "httpd",
		Desc: "HTTP Server",
	}
	err = _db.InsertService(s)
	if err != nil {
		return errors.New(
			"can't insert service: " +
				err.Error(),
		)
	}

	up := db.Metric{
		Service: s.Name,
		Name:    "up",
		Desc:    "true if something is running on httpd port, false otherwise",
	}
	err = _db.InsertMetric(
		up,
	)
	if err != nil {
		return errors.New(
			"can't insert metric: " +
				err.Error(),
		)
	}

	response := db.Metric{
		Service: s.Name,
		Name:    "response",
		Desc:    "true if response is correct",
	}
	err = _db.InsertMetric(
		response,
	)
	if err != nil {
		return errors.New(
			"can't insert metric: " +
				err.Error(),
		)
	}

	v := db.Value{
		Service: s.Name,
		Metric:  up.Name,
		Name:    "up",
		Desc:    "something is running on httpd port",
		Faulty:  false,
	}
	err = _db.InsertValue(
		v,
	)
	if err != nil {
		return errors.New(
			"can't insert value: " +
				err.Error(),
		)
	}

	v = db.Value{
		Service: s.Name,
		Metric:  up.Name,
		Name:    "down",
		Desc:    "nothing is running on httpd port",
		Faulty:  true,
	}
	err = _db.InsertValue(
		v,
	)
	if err != nil {
		return errors.New(
			"can't insert value: " +
				err.Error(),
		)
	}

	p := db.Poller{
		Service: s.Name,
		Metric:  up.Name,
	}
	err = _db.InsertPoller(
		p,
	)
	if err != nil {
		return errors.New(
			"can't insert poller: " +
				err.Error(),
		)
	}

	v = db.Value{
		Service: s.Name,
		Metric:  response.Name,
		Name:    "valid",
		Desc:    "httpd returned valid response",
		Faulty:  false,
	}
	err = _db.InsertValue(
		v,
	)
	if err != nil {
		return errors.New(
			"can't insert value: " +
				err.Error(),
		)
	}

	v = db.Value{
		Service: s.Name,
		Metric:  response.Name,
		Name:    "invalid",
		Desc:    "httpd returned invalid response",
		Faulty:  true,
	}
	err = _db.InsertValue(
		v,
	)
	if err != nil {
		return errors.New(
			"can't insert value: " +
				err.Error(),
		)
	}

	v = db.Value{
		Service: s.Name,
		Metric:  response.Name,
		Name:    "error",
		Desc:    "httpd returned error",
		Faulty:  true,
	}
	err = _db.InsertValue(
		v,
	)
	if err != nil {
		return errors.New(
			"can't insert value: " +
				err.Error(),
		)
	}

	p = db.Poller{
		Service: s.Name,
		Metric:  response.Name,
	}
	err = _db.InsertPoller(
		p,
	)
	if err != nil {
		return errors.New(
			"can't insert poller: " +
				err.Error(),
		)
	}

	messenger := db.Messenger{
		Name: "telegram",
		Desc: "Instant messaging platform",
	}
	err = _db.InsertMessenger(
		messenger,
	)
	if err != nil {
		return errors.New(
			"can't insert messenger: " +
				err.Error(),
		)
	}

	user := db.User{
		Username: "alexandr-bakin",
	}
	err = _db.InsertUser(
		user,
	)
	if err != nil {
		return errors.New(
			"can't insert user: " +
				err.Error(),
		)
	}

	endpoint := db.Endpoint{
		Messenger: "telegram",
		Id:        "456603973",
		User:      user.Username,
		Desc:      "Alexandr Bakin Telegram",
	}
	err = _db.InsertEndpoint(
		endpoint,
	)
	if err != nil {
		return errors.New(
			"can't insert endpoint: " +
				err.Error(),
		)
	}

	n := db.Notifier{
		Messenger: messenger.Name,
	}
	err = _db.InsertNotifier(
		n,
	)
	if err != nil {
		return errors.New(
			"can't insert notifier: " +
				err.Error(),
		)
	}

	t := db.Trigger{
		Service:   "httpd",
		Metric:    "up",
		Value:     "down",
		Messenger: "telegram",
		User:      "alexandr-bakin",
	}

	err = _db.InsertTrigger(t)
	if err != nil {
		return errors.New(
			"can't insert trigger: " +
				err.Error(),
		)
	}

	return nil
}
