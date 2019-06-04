package Route

import (
	"github.com/buguang01/LogService/Service"

	"github.com/buguang01/gsframe/event"
	"github.com/buguang01/util"
)

type LogicRoute struct {
	Msg *event.NsqdMessage
	event.NsqdHander
}

//所在协程的KEY
func (this *LogicRoute) KeyID() string {
	sid := util.NewStringAny(this.Msg.SendSID).ToIntV()
	if Service.Sconf.LogicThreadNum == 0 {
		return this.Msg.SendSID
	}
	return util.NewStringInt(sid % Service.Sconf.LogicThreadNum).ToString()
}

//调用方法
func (this *LogicRoute) Run() {
	this.NsqdHander(this.Msg)
}

func NewLogicRoute(hander event.NsqdHander, msg *event.NsqdMessage) *LogicRoute {
	result := new(LogicRoute)
	result.Msg = msg
	result.NsqdHander = hander
	return result
}
