package service

import (
	"github.com/HaleyLeoZhang/email_server/conf"
	"github.com/HaleyLeoZhang/email_server/constant"
)

// ----------------------------------------------------------------------
// MQ接口限定
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Queue interface {
	Push(s *Service, payload []byte) error
	Pull(s *Service, callback func([]byte) error) error
}

/**
 * 简单工厂
 *
 * @return Queue
 */
func (s *Service) GetEmailQueue() Queue {
	switch conf.Conf.Email.Driver {
	case "amqp":
		return &QueueEngineRabbitMq{}
	case "kafka":
		return &QueueEngineKafka{}
	default:
		panic("email 队列驱动配置选择错误")
	}
}

// ----------------------------------------------------------------------
// RabbitMQ
// ----------------------------------------------------------------------
// 文档 https://github.com/streadway/amqp/tree/master/_examples
// ----------------------------------------------------------------------

type QueueEngineRabbitMq struct{}

func (q *QueueEngineRabbitMq) Push(s *Service, payload []byte) error {
	return s.RabbitMQ.Push(constant.RABBIT_MQ_EXCHANGE, constant.RABBIT_MQ_ROUTIE_KEY, payload)
}

func (q *QueueEngineRabbitMq) Pull(s *Service, callback func([]byte) error) error {
	return s.RabbitMQ.Pull(callback)
}

// ----------------------------------------------------------------------
// Kafka
// ----------------------------------------------------------------------

type QueueEngineKafka struct{}

func (q *QueueEngineKafka) Push(s *Service, payload []byte) error {
	return nil
}

func (q *QueueEngineKafka) Pull(s *Service, callback func([]byte) error) error {
	return nil
}
