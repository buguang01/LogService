package Dal

import (
	"database/sql"
	"fmt"
	"time"
)

//日志本身

type LogInfoMD struct {
	TopicID1  int       `json:"Topic1"` //主题1
	TopicID2  int       `json:"Topic2"` //主题2
	TopicID3  int       `json:"Topic3"` //主题3
	MemberID  int       `json:"MID"`    //用户ID
	ServiceID int       `json:"SID"`    //服务器ID
	UpTime    time.Time //写入时间
	UpNum     int64     //更新数量
	Total     int64     //更新后总数量
	Datas     string    //其他数据
}

func (this *LogInfoMD) GetDB() string {
	return fmt.Sprintf("gamelog_db_%d", this.UpTime.Year())
}

func (this *LogInfoMD) GetTableName() string {
	return fmt.Sprintf("loginfo_%02d_%02d", this.UpTime.Month(), (this.UpTime.Day()-1)/3+1)
}

func InsertTopic(conndb *sql.DB, md *LogInfoMD) error {
	sqlstr := `
	INSERT INTO %s.%s(
		topicid1,topicid2,topicid3,memberid, serviceid,uptime,upnum,total, datas
	)VALUES(
		?,?,?,?, ?,?,?,?, ?
	)	;
	`
	sqlstr = fmt.Sprintf(sqlstr, md.GetDB(), md.GetTableName())
	_, err := conndb.Exec(sqlstr,
		md.TopicID1, md.TopicID2, md.TopicID3, md.MemberID, md.ServiceID, md.UpTime, md.UpNum, md.Total, md.Datas)
	if err != nil {
		return err
	}
	return nil
}
