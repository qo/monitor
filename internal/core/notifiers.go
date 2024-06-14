package core

import (
	"github.com/qo/monitor/examples/plugins/notifiers/telegram"
	"github.com/qo/monitor/internal/plugins"
)

// Список публикующих плагинов
var notifierPlugins = []plugins.Notifier{
	telegram.Notifier,
}
