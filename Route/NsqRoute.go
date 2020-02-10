package Route

import (
	"LogService/ActionCode"
	"LogService/Events/NsqEvents"

	"github.com/buguang01/bige/messages"
)

func init() {
	NsqRoute = messages.JsonMessageHandleNew()
	NsqRoute.SetRoute(ActionCode.Log_Send, &NsqEvents.LogEvent{})
}

var (
	NsqRoute messages.IMessageHandle
)
