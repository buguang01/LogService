package Service

import (
	"github.com/buguang01/gsframe/model"
	"github.com/buguang01/gsframe/module"
	"github.com/buguang01/gsframe/runserver"
)

var (
	GameExample  *runserver.GameServiceBase
	DBExample    *module.SqlDataModule
	LogicExample *module.LogicModule
	MysqlExample *model.MysqlAccess
	NsqdExample  *module.NsqdModule
)

func ServiceStop() {
	NsqdExample.StopConsumer()
}
