package main

import (
	"LogService/Flag"
	"LogService/Manage"
	"LogService/Route"
	"LogService/Service"
	"flag"
	"io/ioutil"
	"os"
	"sync"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/json"
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/module"
	"github.com/buguang01/bige/runserver"
	_ "github.com/icattlecoder/godaemon"
)

func main() {
	Service.Sconf = new(Service.ServiceConf)
	var conf = Service.Sconf
	if !flag.Parsed() {
		flag.Parse()
	}
	f, err := os.Open(*Flag.Flagc)
	if err != nil {
		panic(err)
	}
	b, _ := ioutil.ReadAll(f)
	f.Close()
	json.Unmarshal(b, &conf)
	Logger.Init(conf.LogLv, conf.LogPath, conf.LogMode)
	defer Logger.LogClose()

	//启动服务
	Service.GameExample = runserver.NewGameService(&conf.GameConf)
	Service.GameExample.ServiceStopHander = Service.ServiceStop
	Service.MysqlExample = model.NewMysqlAccess(&conf.DBConf)
	if err := Service.MysqlExample.Ping(); err != nil {
		Logger.PFatal(err)
		return
	}
	defer Service.MysqlExample.Close()
	Service.LogicExample = module.NewLogicModule(&conf.LogicConf)
	// Service.DBExample = module.NewSqlDataModule(&conf.SqlConf, Service.MysqlExample.GetDB())
	Service.NsqdExample = module.NewNsqdModule(&conf.NsqdConf, conf.GameConf.ServiceID)

	Service.NsqdExample.RouteFun = Route.NsqRouteHander
	// Service.GameExample.AddModule(Service.DBExample)
	Service.GameExample.AddModule(Service.NsqdExample)
	Service.GameExample.AddModule(Service.LogicExample)
	InitData()
	Service.GameExample.Run()
}

func InitData() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	Manage.TopicManageEx.Load(wg)
	wg.Wait()
}
