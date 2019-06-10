package Manage

import (
	"fmt"
	"sync"
	"time"

	"github.com/buguang01/LogService/Dal"
	"github.com/buguang01/LogService/Service"

	"github.com/buguang01/Logger"
	"github.com/buguang01/util"
)

type TopicManage struct {
	topiclist map[string]int
	maplock   sync.RWMutex           //锁
	topicID   int                    //新主题的ID
	dblist    map[string]*LogDbModel //数据库
	dblock    sync.RWMutex           //锁
}

func init() {
	TopicManageEx = new(TopicManage)
	TopicManageEx.topiclist = make(map[string]int)
	TopicManageEx.dblist = make(map[string]*LogDbModel)
}

//主题管理器
var TopicManageEx *TopicManage

func (this *TopicManage) Load(wg *sync.WaitGroup) {
	defer wg.Done()
	rows := Dal.SelectTopicInfoByAll()
	defer rows.Close()
	for rows.Next() {
		id := 0
		name := ""
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		this.topiclist[name] = id
	}
	this.topicID = Dal.GetTopicMax()
}

//GetTopicID拿到主题的ID
func (this *TopicManage) GetTopicID(name string) int {
	if name == "" {
		return -1
	}
	this.maplock.RLock()
	result, ok := this.topiclist[name]
	this.maplock.RUnlock()
	if ok {
		return result
	}
	//没找到就要写入了
	this.maplock.Lock()

	result, ok = this.topiclist[name]
	if ok {
		this.maplock.Unlock()
		return result
	} else {
		this.topicID++
		topid := this.topicID
		this.topiclist[name] = topid
		result = topid
		this.maplock.Unlock()
		//要写入数据库
		err := Dal.UpTopicInfo(&Dal.TopicInfoMD{UID: topid, Name: name})
		if err != nil {
			Logger.PError(err, "UpTopicInfo:%d,%s", topid, name)
		}
	}
	return result
}

func (this *TopicManage) SendDB(logmd *Dal.LogInfoMD) {
	var result *LogDbModel
	var ok bool
	util.UsingRead(&this.dblock, func() {
		result, ok = this.dblist[logmd.GetDB()]
	})
	if !ok {
		util.UsingWiter(&this.dblock, func() {
			result, ok = this.dblist[logmd.GetDB()]
			if ok {
				return
			}
			result = new(LogDbModel)
			result.Name = logmd.GetDB()
			result.CreateTime = logmd.UpTime
			this.dblist[result.Name] = result
			result.CreateDB()
		})
	}
	err := Dal.InsertTopic(Service.MysqlExample.DBConobj, logmd)
	if err != nil {
		panic(err)
	}
	err = Dal.UpTopicTree(Service.MysqlExample.DBConobj, logmd)
	if err != nil {
		panic(err)
	}

}

type LogDbModel struct {
	Name       string    //数据库名字
	CreateTime time.Time //生成时间点
	//一年一库
}

func (this *LogDbModel) CreateDB() {
	sqlstr := `
	CREATE DATABASE IF NOT EXISTS gamelog_db_%d  DEFAULT CHARACTER SET utf8mb4 ;	
	`
	conndb := Service.MysqlExample.DBConobj
	conndb.Exec(fmt.Sprintf(sqlstr, this.CreateTime.Year()))
	tabstr := `
	CREATE TABLE IF NOT EXISTS  gamelog_db_%d.loginfo_%02d_%02d (
		uid bigint(20) NOT NULL AUTO_INCREMENT,
		topicid1 int(11) NOT NULL,
		topicid2 int(11) NOT NULL,
		topicid3 int(11) NOT NULL,
		memberid int(11) NOT NULL,
		serviceid int(11) NOT NULL,
		uptime datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
		upnum bigint(20) NOT NULL,
		total bigint(20) NOT NULL,
		datas varchar(4000) NOT NULL,
		PRIMARY KEY (uid),
		KEY topicid1 (memberid,topicid1,topicid2,topicid3),
		KEY uptime (uptime)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`
	for i := 1; i <= 12; i++ {
		for r := 1; r <= 11; r++ {
			conndb.Exec(fmt.Sprintf(tabstr, this.CreateTime.Year(), i, r))
		}
	}
}
