package queue

// ----------------------------------------------------------------------
// RabbitMQ
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
// 文档 https://github.com/streadway/amqp/tree/master/_examples
// ----------------------------------------------------------------------

import (
	"email_server/pkg/setting"
	"fmt"
	"github.com/streadway/amqp"
	// "time"
)

type AMQP struct {
	// 公共
	Payload string
	// 当前驱动配置项
	Exchange string
	Queue    string
}

func (a *AMQP) newConnect() (*amqp.Connection, error) {
	dial := fmt.Sprintf("amqp://%v:%v@%v:%v/", setting.AMQPSetting.USER, setting.AMQPSetting.PASSWORD, setting.AMQPSetting.HOST, setting.AMQPSetting.PORT)
	return amqp.Dial(dial)
}

func (a *AMQP) Push() error {

	conn, err := a.newConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.Publish(
		a.Exchange, // exchange
		a.Queue,    // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(a.Payload),
		})

	if err != nil {
		return err
	}

	conn.Close()
	return nil
}

func (a *AMQP) Pull(callback func(string) error) error {
	// 最大协程数量
	pool := make(chan int, setting.AMQPSetting.CHANNEL_NUMBER)
	defer close(pool)

	conn, err := a.newConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	delivery, err := ch.Consume(
		a.Queue, // name
		"",      // consumerTag,
		false,   // noAck
		false,   // exclusive
		false,   // noLocal
		false,   // noWait
		nil,     // arguments
	)

	for {
		select {
		case d := <-delivery:
			pool <- 1
			go a.handle(d, callback, pool)
		}
	}
	return nil
}

func (a *AMQP) handle(d amqp.Delivery, callback func(string) error, pool <-chan int) error {

	err := callback(string(d.Body))
	if err != nil {
		return err
	}
	d.Ack(false)
	<-pool
	return nil
}
