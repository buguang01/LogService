package NsqEvents

import (
	"LogService/Events"
	"LogService/Service"

	"github.com/buguang01/bige/messages"
)

//日志逻辑
type LogEvent struct {
	messages.NsqdMessage
	messages.LogicMessage
	Data Events.LogLogic
}

func (msg *LogEvent) NsqDirectCall() {
	//设置Logic层，对应的协程用户ID
	msg.UserID = msg.Data.MID
	Service.LogicExample.AddMsg(msg)
}

func (msg *LogEvent) MessageHandle() {
	msg.Data.Log_Send()

}
