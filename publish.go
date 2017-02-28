package goRabbitMq

import (
	"errors"
	"fmt"
	"goRabbitMq/common"
	"goRabbitMq/configuration"

	"time"

	"github.com/streadway/amqp"
)

func (handle *MessageHandle) Publish(msgPackage *MessagePackage) PublishResponse {
	defer func() {
		if err := recover(); err != nil {
			common.Log.Errorf("publish message to rabbitmq error, %s", err)
		}
	}()
	execute := func(msgPackage *MessagePackage) PublishResponse {
		//获取配置
		cfg := configuration.ReadCfg(msgPackage.AppId, msgPackage.Code, &configuration.Cfg{})
		if cfg == (&configuration.Cfg{}) {
			common.Log.Errorf("can't cfg file")
			return PublishResponse{Code: 201, Msg: "can't cfg file"}
		}
		//common.Log.Info("msg :", msgPackage)
		//common.Log.Info("cfg :", cfg)
		//序列化消息
		by := JsonToBytes(msgPackage.Body)
		//发布
		now := time.Now()
		err := publishmessage(by, msgPackage.MsgUniqueId, cfg)
		if err != nil {
			panic(err)
			return PublishResponse{Code: 500, Msg: err.Error()}
		}
		runTime := time.Now().Sub(now).Seconds() * 1000

		m := ToString("run ok ,milliseconds  ", runTime)
		return PublishResponse{Code: 200, Msg: m}
	}
	return execute(msgPackage)
}

//发布消息
func publishmessage(body []byte, messageId string, cfg *configuration.Cfg) error {
	//common.Log.Info("start publish message")
	//channel

	con, err := PoolInit()
	if err != nil {
		return errors.New("PoolInit error")
	}
	channel, err := con.GetChannel(cfg.AppId, cfg.Code)
	if err != nil {
		common.Log.Errorf("get channel error %s", err)
		return err
	}
	//exchange declared
	ExchangeDeclare(cfg.Exchange, cfg.ExchangeType, channel)
	if cfg.Reliable {
		common.Log.Info("enabling publishing confirms.")
		if err := channel.Confirm(false); err != nil {
			common.Log.Errorf("Channel could not be put into confirm mode: %s", err)
			return err
		}
		confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))
		defer confirmOne(confirms)
	}
	//common.Log.Infof("declared Exchange, publishing %dB body (%q)", len(body), body)
	handMap := make(map[string]interface{}, 1)
	handMap["msgid"] = messageId
	handMap["uuid"] = "1"
	if err = channel.Publish(
		cfg.Exchange,   // publish to an exchange
		cfg.RoutingKey, // routing to 0 or more queues
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			Headers:         handMap,
			ContentType:     "application/json",
			ContentEncoding: cfg.Encoding,
			Body:            body,
			DeliveryMode:    If(cfg.MsgPersistent, amqp.Persistent, amqp.Transient).(uint8), // 1=non-persistent, 2=persistent
			Priority:        0,                                                              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}

	return nil
}

//发布消息确认
func confirmOne(confirms <-chan amqp.Confirmation) {
	common.Log.Infof("waiting for confirmation of one publishing")
	if confirmed := <-confirms; confirmed.Ack {
		common.Log.Infof("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		common.Log.Infof("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}
