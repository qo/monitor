package plugins

import "github.com/qo/monitor/internal/db"

// Вообще говоря, следует использовать
// интерфейс с необходимыми методами
// для работы с БД
// вместо того, чтобы импортировать пакет.
// Но для простоты просто импортируем пакет

// Интерфейс для сущности
// опрашивающий плагин
type Poller interface {

	// Получить имя сервиса.
	// Данный метод предназначен для того,
	// чтобы связать сущность poller
	// из БД с кодом плагина.
	Service() string

	// Получить имя метрики.
	// Предназначение аналогично методу
	// Service
	Metric() string

	// Получить текущее значение
	// данной метрики данного сервиса
	Value() (string, error)
}

// Интерфейс для сущности
// публикующий плагин
type Notifier interface {

	// Получить имя мессенджера.
	// Данный метод предназначен для того,
	// чтобы связать сущность notifier
	// из БД с кодом плагина
	Messenger() string

	// Отослать сообщение
	// о данном значении
	// в данный мессенджер
	// на данный эндпоинт
	Notify(
		value db.Value,
		endpoint db.Endpoint,
	) error

	// Данный метод предназначен для того,
	// чтобы предоставить разработчику
	// возможность зарегистрировать
	// обработчики для того,
	// чтобы пользователи могли оставлять
	// обратную связь по задачам
	Run()
}
