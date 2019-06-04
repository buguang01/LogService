package Service

import (
	"github.com/buguang01/Logger"
	"github.com/buguang01/gsframe/model"
	"github.com/buguang01/gsframe/module"
	"github.com/buguang01/gsframe/runserver"
)

type ServiceConf struct {
	GameConf       runserver.GameConfigModel
	LogicConf      module.LogicConfig
	SqlConf        module.SqlDataConfig
	NsqdConf       module.NsqdConfig
	DBConf         model.MysqlConfigModel
	LogLv          Logger.LogLevel
	LogPath        string
	SqlThreadNum   int
	LogicThreadNum int
}

var Sconf *ServiceConf
