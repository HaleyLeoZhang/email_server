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
	"sync"
)

type AMQP struct {
	// 公共
	Payload []byte
	// 当前驱动配置项
	Exchange string
	Queue    string
}

/**
 * 单例模式
 */
type oneAMQP struct {
	Conn *amqp.Connection
	// 因为多协程共用一个tcp链接,防止并发交错错写入
	// 但一个链接能建立多个通道
	Lock sync.Mutex
}

var one oneAMQP

func (a *AMQP) newConnect() *amqp.Connection {
	if nil == one.Conn {
		dial := fmt.Sprintf("amqp://%v:%v@%v:%v/", setting.AMQPSetting.USER, setting.AMQPSetting.PASSWORD, setting.AMQPSetting.HOST, setting.AMQPSetting.PORT)
		var err error
		one.Conn, err = amqp.Dial(dial)
		if err != nil {
			panic([]string{err.Error()})
		}
	}
	return one.Conn
}

func (a *AMQP) SetPayload(payload []byte) {
	a.Payload = payload
}

func (a *AMQP) Push() error {

	conn := a.newConnect()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	one.Lock.Lock()
	err = ch.Publish(
		a.Exchange, // exchange
		a.Queue,    // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        a.Payload,
		})
	one.Lock.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func (a *AMQP) Pull(callback func(string) error) error {

	conn := a.newConnect()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	one.Lock.Lock()
	delivery, err := ch.Consume(
		a.Queue, // name
		"",      // consumerTag,
		false,   // noAck
		false,   // exclusive
		false,   // noLocal
		false,   // noWait
		nil,     // arguments
	)
	one.Lock.Unlock()

	for {
		select {
		case d := <-delivery:
			pool <- 1
			go a.handle(d, callback)
		}
	}
	return nil
}

func (a *AMQP) handle(d amqp.Delivery, callback func(string) error) error {

	err := callback(string(d.Body))
	if err != nil {
		return err
	}
	d.Ack(false)
	<-pool
	return nil
}
