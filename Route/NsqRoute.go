package Route

import (
	"LogService/ActionCode"
	"LogService/Events"
	"LogService/Service"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
)

func init() {
	NsqRouteList = make(map[int]event.NsqdHander)
	NsqRouteList[ActionCode.Log_Send] = Events.EventLog
}

var (
	NsqRouteList map[int]event.NsqdHander
)

func NsqRouteHander(msg event.INsqdMessage) {
	//这里就要写消息的确认处理方法
	//和消息处理方法的运行
	hander, ok := NsqRouteList[msg.GetActionID()]
	if ok {
		//把运行逻辑放到按服务器来源分的协程中
		logicmd := NewLogicRoute(hander, msg)
		Service.LogicExample.AddMsg(logicmd)
		//hander(msg)
	} else {
		Logger.PError(nil, "Nsq Action:%d,not hander.", msg.GetActionID())
	}
}
