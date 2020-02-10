package NsqEvents

import (
	"LogService/Events"
	"LogService/Service"

	"github.com/buguang01/bige/messages"
)

//日志逻辑
type LogEvent struct {
	messages.NsqdMessage
	Events.LogLogic
}

func (msg *LogEvent) NsqDirectCall() {
	//设置Logic层，对应的协程用户ID
	msg.UserID = msg.MID
	Service.LogicExample.AddMsg(msg)
}
