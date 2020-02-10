package Service

import (
	"time"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/model"
)

type ServiceConf struct {
	ServiceID       int           //游戏服务器ID
	PStatusTime     time.Duration //打印状态的时间（秒）
	NsqdAddr        []string
	NsqdLookupdAddr []string
	DBConf          model.MysqlConfigModel
	LogLv           Logger.LogLevel
	LogPath         string
	LogMode         Logger.LogMode
}

var Sconf *ServiceConf
