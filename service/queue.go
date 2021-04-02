package service

import (
	"context"
	"github.com/HaleyLeoZhang/email_server/conf"
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/go-component/driver/xkafka"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/errgroup"
	"github.com/Shopify/sarama"
	"strings"
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
	case constant.DRIVER_NAME_RABIT_MQ:
		return &QueueEngineRabbitMq{}
	case constant.DRIVER_NAME_KAFKA:
		return &QueueEngineKafka{}
	default:
		panic("vo 队列驱动配置错误")
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
	s.KafkaEmailProducer.SendMsgAsync(conf.Conf.KafkaTopic.TopicList[0], payload)
	return nil
}

func (q *QueueEngineKafka) Pull(s *Service, callback func([]byte) error) error {
	registerHandler := func(session *xkafka.ConsumerSession, msgs <-chan *sarama.ConsumerMessage) (errKafka error) {
		fun := xkafka.IteratorBatchFetch(session, msgs, conf.Conf.Email.BatchNumber, constant.KAFKA_BATCH_WAIT_MAX_SECOND)
		for {
			kafkaMessages, ok := fun()
			if !ok {
				return
			}
			eg := &errgroup.Group{}
			eg.GOMAXPROCS(conf.Conf.Email.Consumer)
			for _, business := range kafkaMessages {
				tmp := business
				eg.Go(func(context.Context) error {
					err := callback(tmp.Value)
					if err != nil {
						xlog.Errorf("Kafka Callback Value(%v) Group(%+v) Topic(%v) Err(%v)",
							string(tmp.Value), conf.Conf.KafkaTopic.Group,
							strings.Join(conf.Conf.KafkaTopic.TopicList, ","), err)
					}else{
						xlog.Infof("Kafka Callback Value(%v) Group(%+v) Topic(%v) Success",
							string(tmp.Value), conf.Conf.KafkaTopic.Group,
							strings.Join(conf.Conf.KafkaTopic.TopicList, ","))
					}
					return nil
				})
			}
			_ = eg.Wait()
		}
	}
	s.KafkaEmailConsumer.Consumer.RegisterHandler(registerHandler)
	s.KafkaEmailConsumer.Start()
	return nil
}
