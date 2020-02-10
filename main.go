package main

import (
	"LogService/Flag"
	"LogService/Manage"
	"LogService/Route"
	"LogService/Service"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/modules"
	"github.com/buguang01/util"
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
	Service.GameExample = modules.NewGameService(
		modules.GameServiceSetSID(Service.Sconf.ServiceID),
		modules.GameServiceSetPTime(Service.Sconf.PStatusTime),
	)
	Service.MysqlExample = model.NewMysqlAccess(&conf.DBConf)
	if err := Service.MysqlExample.Ping(); err != nil {
		Logger.PFatal(err)
		return
	}
	defer Service.MysqlExample.Close()

	Service.LogicExample = modules.NewLogicModule()
	Service.NsqdExample = modules.NewNsqdModule(
		modules.NsqdSetPorts(conf.NsqdAddr...),
		modules.NsqdSetLookup(conf.NsqdLookupdAddr...),
		modules.NsqdSetMyTopic(util.ToString(Service.GameExample.ServiceID)),
		modules.NsqdSetMyChannelName(fmt.Sprintf("chancel_%d", Service.GameExample.ServiceID)),
		modules.NsqdSetRoute(Route.NsqRoute),
	)
	Service.GameExample.AddModule(
		Service.LogicExample,
		Service.NsqdExample)
	InitData()
	Service.GameExample.Run()
}

func InitData() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	Manage.TopicManageEx.Load(wg)
	wg.Wait()
}
