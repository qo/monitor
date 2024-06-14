package core

import (
	"github.com/qo/monitor/examples/plugins/notifiers/telegram"
	"github.com/qo/monitor/internal/plugins"
)

// Публикующие плагины
var notifierPlugins = []plugins.Notifier{
	telegram.Notifier,
}
