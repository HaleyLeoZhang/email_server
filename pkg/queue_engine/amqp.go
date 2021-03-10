package queue_engine

// ----------------------------------------------------------------------
// RabbitMQ
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
// 文档 https://github.com/streadway/amqp/tree/master/_examples
// ----------------------------------------------------------------------

import (
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/email_server/service"
)

type RabbitMq struct{}

func (a *RabbitMq) Push(s *service.Service, payload []byte) error {
	return s.RabbitMQ.Push(constant.RABBIT_MQ_EXCHANGE, constant.RABBIT_MQ_ROUTIE_KEY, payload)
}

func (a *RabbitMq) Pull(s *service.Service, callback func([]byte) error) error {
	return s.RabbitMQ.Pull(callback)
}
