package Dal

import (
	"LogService/Service"
	"database/sql"
)

//日志分类主题
//TopicInfoMD
type TopicInfoMD struct {
	UID  int    //日志ID
	Name string //日志事件名字
}

func SelectTopicInfoByAll() *sql.Rows {
	sqlstr := `
	SELECT uid,name FROM topicinfo;
	`
	result, err := Service.MysqlExample.DBConobj.Query(sqlstr)
	if err != nil {
		panic(err)
	}
	return result
}

//GetTopicMax拿日志主题最大ID
func GetTopicMax() int {
	result := 0
	sqlstr := `
	SELECT IFNULL(Max(uid),0) FROM topicinfo;
	`
	row := Service.MysqlExample.DBConobj.QueryRow(sqlstr)

	err := row.Scan(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func UpTopicInfo(md *TopicInfoMD) error {
	sqlstr := `
	INSERT INTO topicinfo(
		uid,name
	)VALUES(
		?,?
	);
	`
	_, err := Service.MysqlExample.DBConobj.Exec(sqlstr, md.UID, md.Name)
	if err != nil {
		return err
	}
	return nil
}
