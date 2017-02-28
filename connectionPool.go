package goRabbitMq

import (
	"errors"
	"goRabbitMq/common"
	"goRabbitMq/configuration"

	"github.com/streadway/amqp"
)

type Conn struct {
}

var index int
var connPool map[string][]*amqp.Connection
var isinit bool

//pool init
func PoolInit() (*Conn, error) {
	if isinit {
		common.Log.Infof("is init true")
		return &Conn{}, nil
	}
	cfgs := configuration.AllCfgs()
	if len(cfgs) <= 0 {
		common.Log.Errorf("can't load cfg , init pool fail")
		return &Conn{}, errors.New("can't load cfg , init pool fail")
	}
	//****注意写法，否则掉坑
	connPool = make(map[string][]*amqp.Connection)
	for _, k := range cfgs {
		//common.Log.Infof("appid %s,code %s create conn,pool size %d", k.AppId, k.Code, k.MaxConn)
		list := make([]*amqp.Connection, k.MaxConn)
		key := ToString(k.AppId, "_", k.Code)
		for i := 0; i < k.MaxConn; i++ {
			con, err := newConn(k.Host)
			if err != nil {
				common.Log.Errorf("create conn error: %s", err)
			} else {
				list[i] = con
			}
		}
		//common.Log.Infof("appid %s,code %s create conn,pool key %s,pool size %d", k.AppId, k.Code, key, cap(list))
		connPool[key] = list
	}
	//common.Log.Infof("pool init ok ,pool size %d", len(connPool))
	isinit = true
	return &Conn{}, nil
}

//get channel
func (cn *Conn) GetChannel(appId string, code string) (*amqp.Channel, error) {
	common.Log.Infof("connPool =%d", len(connPool))
	for k, v := range connPool {
		common.Log.Infof("k=%s,v=%d", k, len(v))
	}
	key := ToString(appId, "_", code)
	list := connPool[key]
	//common.Log.Infof("appid %s,code %s,get conn index %d, key %s,pool size %d", appId, code, index, key, len(list))
	i := index % len(list)
	index = index + 1
	conn := list[i]
	//common.Log.Info("channel create ok", index)
	return newChannel(conn)
}

func newConn(amqpURI string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(amqpURI)
	if err != nil {
		common.Log.Errorf("Dial: %s", err)
		return &amqp.Connection{}, err
	}
	return connection, nil
}

func newChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		common.Log.Errorf("Channel: %s", err)
	}
	return ch, nil
}
