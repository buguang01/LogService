package Route

import (
	"LogService/Service"

	"github.com/buguang01/bige/event"
	"github.com/buguang01/util"
)

type LogicRoute struct {
	Msg event.INsqdMessage
	event.NsqdHander
}

//所在协程的KEY
func (this *LogicRoute) KeyID() string {
	sid := util.NewStringAny(this.Msg.GetSendSID()).ToIntV()
	if Service.Sconf.LogicConf.InitNum == 0 {
		return this.Msg.GetSendSID()
	}
	return util.NewStringInt(sid % Service.Sconf.LogicConf.InitNum).ToString()
}

//调用方法
func (this *LogicRoute) Run() {
	this.NsqdHander(this.Msg)
}

func NewLogicRoute(hander event.NsqdHander, msg event.INsqdMessage) *LogicRoute {
	result := new(LogicRoute)
	result.Msg = msg
	result.NsqdHander = hander
	return result
}
