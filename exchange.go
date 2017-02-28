package goRabbitMq

import (
	"goRabbitMq/common"

	"github.com/streadway/amqp"
)

var exDeclared bool

func ExchangeDeclare(name, exType string, channel *amqp.Channel) {
	if exDeclared {
		return
	}
	if err := channel.ExchangeDeclare(
		name,   // name
		exType, // type
		true,   // durable
		false,  // auto-deleted
		false,  // internal
		false,  // noWait
		nil,    // arguments
	); err != nil {
		common.Log.Errorf("ExchangeDeclare error: %s", err)
	}
	//common.Log.Infof("got Channel, declaring %q Exchange (%q)", name, exType)
	exDeclared = true
}
