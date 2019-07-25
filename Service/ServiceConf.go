package Service

import (
	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/module"
	"github.com/buguang01/bige/runserver"
)

type ServiceConf struct {
	GameConf  runserver.GameConfigModel
	LogicConf module.LogicConfig
	SqlConf   module.SqlDataConfig
	NsqdConf  module.NsqdConfig
	DBConf    model.MysqlConfigModel
	LogLv     Logger.LogLevel
	LogPath   string
	LogMode   Logger.LogMode
}

var Sconf *ServiceConf
