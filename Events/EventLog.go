package Events

import (
	"LogService/Dal"
	"LogService/Manage"
	"time"

	"github.com/buguang01/bige/event"
	"github.com/buguang01/util"
)

//日志逻辑
//EventLog
func EventLog(msg event.INsqdMessage) {
	//收到的日志
	et := event.JsonMap(msg.GetData().(map[string]interface{}))
	logmd := new(Dal.LogInfoMD)
	logmd.ServiceID = util.NewStringAny(et["SID"]).ToIntV()
	logmd.MemberID = util.NewStringAny(et["MID"]).ToIntV()
	logmd.TopicID1 = Manage.TopicManageEx.GetTopicID(util.NewStringAny(et["Topic1"]).ToString())
	logmd.TopicID2 = Manage.TopicManageEx.GetTopicID(util.NewStringAny(et["Topic2"]).ToString())
	logmd.TopicID3 = Manage.TopicManageEx.GetTopicID(util.NewStringAny(et["Topic3"]).ToString())
	logmd.UpTime, _ = time.Parse(time.RFC3339, util.NewStringAny(et["UpTime"]).ToString())
	logmd.UpNum = util.NewStringAny(et["UpNum"]).ToInt64V()
	logmd.Total = util.NewStringAny(et["Total"]).ToInt64V()
	logmd.Datas = util.NewStringAny(et["Datas"]).ToString()
	Manage.TopicManageEx.SendDB(logmd)
}
