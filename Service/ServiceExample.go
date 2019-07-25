package Service

import (
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/module"
	"github.com/buguang01/bige/runserver"
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
