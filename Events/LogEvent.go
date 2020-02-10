package Events

import (
	"LogService/Dal"
	"LogService/Manage"
	"time"

	"github.com/buguang01/bige/messages"
)

//日志逻辑
type LogLogic struct {
	messages.LogicMessage
	Topic1 string    `json:"Topic1"` //主题1
	Topic2 string    `json:"Topic2"` //主题2
	Topic3 string    `json:"Topic3"` //主题3
	MID    int       `json:"MID"`    //用户ID
	SID    int       `json:"SID"`    //服务器ID
	UpTime time.Time //写入时间
	UpNum  int64     //更新数量
	Total  int64     //更新后总数量
	Datas  string    //其他数据
}

//调用方法
func (msg *LogLogic) MessageHandle() {
	logmd := new(Dal.LogInfoMD)
	logmd.ServiceID = msg.SID
	logmd.MemberID = msg.MID
	logmd.TopicID1 = Manage.TopicManageEx.GetTopicID(msg.Topic1)
	logmd.TopicID2 = Manage.TopicManageEx.GetTopicID(msg.Topic2)
	logmd.TopicID3 = Manage.TopicManageEx.GetTopicID(msg.Topic3)
	logmd.UpTime = msg.UpTime
	logmd.UpNum = msg.UpNum
	logmd.Total = msg.Total
	logmd.Datas = msg.Datas
	Manage.TopicManageEx.SendDB(logmd)
}
