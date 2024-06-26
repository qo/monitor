package core

import (
	"github.com/qo/monitor/examples/plugins/pollers/httpd/response"
	"github.com/qo/monitor/examples/plugins/pollers/httpd/up"
	"github.com/qo/monitor/internal/plugins"
)

// Список подписывающих плагинов
var pollerPlugins = []plugins.Poller{
	httpd_up.Poller,
	httpd_response.Poller,
}
