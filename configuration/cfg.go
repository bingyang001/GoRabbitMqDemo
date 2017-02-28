package configuration

import (
	"goRabbitMq/common"
	"strconv"
	"strings"

	"github.com/Unknwon/goconfig"
)

type Cfg struct {
	AppId         string
	Code          string
	Host          string
	Exchange      string
	ExchangeType  string
	RoutingKey    string
	Encoding      string
	Reliable      bool
	MaxConn       int
	MsgPersistent bool
}

var (
	cfg       *goconfig.ConfigFile
	isload    bool
	emptyKeys []string
	emptyCfgs []*Cfg
)

//获取所有配置
func AllCfgs() []*Cfg {
	allKeys := ReadAllCfgKeys()
	if len(allKeys) == 0 {
		return emptyCfgs
	}
	cfgArr := make([]*Cfg, 1)
	for _, v := range allKeys {
		karray := strings.Split(v, "_")
		cfg := ReadCfg(karray[0], karray[1], &Cfg{})
		if cfg == (&Cfg{}) {
			common.Log.Errorf("can't cfg file")
		} else {
			c := &Cfg{
				Host:          cfg.Host,
				Exchange:      cfg.Exchange,
				ExchangeType:  cfg.ExchangeType,
				RoutingKey:    cfg.RoutingKey,
				Encoding:      cfg.Encoding,
				Reliable:      cfg.Reliable,
				MaxConn:       cfg.MaxConn,
				AppId:         cfg.AppId,
				Code:          cfg.Code,
				MsgPersistent: cfg.MsgPersistent}
			//cfgArr = append(cfgArr, c)
			cfgArr[0] = c
		}
	}
	common.Log.Infof("load cfg item count %d", len(cfgArr))
	return cfgArr
}

//read all cfg keys
func ReadAllCfgKeys() []string {
	cfg, _ := loadCfgfile("./conf.ini")
	return cfg.GetSectionList()
}

//read configuration
func ReadCfg(appid string, code string, defaultValue *Cfg) *Cfg {
	return defaultCfg()
	//	cfgs, err := loadCfgfile("./conf.ini")
	//	if err != nil {
	//		return defaultValue
	//	}
	//	s := []string{appid, code}
	//	var superKey = strings.Join(s, "_")
	//	mapCfg, err := cfgs.GetSection(superKey)
	//	if err != nil {
	//		common.Log.Errorf("can't read cfg %s", err.Error())
	//		return defaultValue
	//	}
	//	maxconn, _ := strconv.Atoi(mapCfg["MaxConn"])
	//	c := &Cfg{
	//		Host:          mapCfg["Host"],
	//		Exchange:      mapCfg["Exchange"],
	//		ExchangeType:  mapCfg["ExchangeType"],
	//		RoutingKey:    mapCfg["RoutingKey"],
	//		Encoding:      mapCfg["Encoding"],
	//		Reliable:      toBool(mapCfg["Reliable"], false),
	//		MaxConn:       maxconn,
	//		AppId:         mapCfg["AppId"],
	//		Code:          mapCfg["Code"],
	//		MsgPersistent: toBool(mapCfg["MsgPersistent"], false)}
	//	return c
}

func loadCfgfile(name string) (*goconfig.ConfigFile, error) {
	if isload == true {
		common.Log.Infof("get cache cfg")
	}
	if isload == false {
		cfgs, err := goconfig.LoadConfigFile("./conf.ini")
		if err != nil {
			common.Log.Errorf("can't load cfg file %s", err.Error())
			return &goconfig.ConfigFile{}, err
		}
		isload = true
		cfg = cfgs
	}
	return cfg, nil
}

func toBool(str string, defaultVal bool) bool {
	val, err := strconv.ParseBool(str)
	if err != nil {
		return defaultVal
	}
	return val
}

func defaultCfg() *Cfg {
	c := &Cfg{
		Host:          "amqp://guest:guest@172.16.100.48:5672/",
		Exchange:      "liguo",
		ExchangeType:  "topic",
		RoutingKey:    "",
		Encoding:      "utf-8",
		Reliable:      false,
		MaxConn:       3,
		AppId:         "test2",
		Code:          "liguo",
		MsgPersistent: false}
	return c
}
