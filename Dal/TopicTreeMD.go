package Dal

import "database/sql"

//日志树

func UpTopicTree(conndb *sql.DB, md *LogInfoMD) error {
	sqlstr := `
	INSERT INTO logtree(
		topicid1,topicid2,topicid3
	)VALUES(
		?,?,?
	)
	ON DUPLICATE KEY UPDATE
	topicid1=values(topicid1)
	;
	`
	_, err := conndb.Exec(sqlstr,
		md.TopicID1, md.TopicID2, md.TopicID3)
	if err != nil {
		return err
	}
	return nil
}
