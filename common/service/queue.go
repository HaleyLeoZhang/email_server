package service

import (
	"context"
	"github.com/HaleyLeoZhang/email_server/common/constant"
	"github.com/HaleyLeoZhang/go-component/driver/xkafka"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/driver/xrabbitmq"
	"github.com/HaleyLeoZhang/go-component/errgroup"
	"github.com/Shopify/sarama"
	"strings"
)

// ----------------------------------------------------------------------
// MQ驱动限定
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Queue interface {
	Start(s *Service)
	Close(s *Service)
	Push(s *Service, payload []byte) error
	Pull(s *Service, callback func([]byte) error) error
}

/**
 * 简单工厂
 *
 * @return Queue
 */
func (s *Service) GetEmailQueue() (q Queue) {
	switch s.cfg.Email.Driver {
	case constant.DRIVER_NAME_RABIT_MQ:
		q = &QueueEngineRabbitMq{}
	case constant.DRIVER_NAME_KAFKA:
		q = &QueueEngineKafka{}
	default:
		panic("vo 队列驱动配置错误")
	}
	// 初始化配置
	q.Start(s)
	return
}

// ----------------------------------------------------------------------
// RabbitMQ
// ----------------------------------------------------------------------
// 文档 https://github.com/streadway/amqp/tree/master/_examples
// ----------------------------------------------------------------------

type QueueEngineRabbitMq struct{}

func (q *QueueEngineRabbitMq) Start(s *Service) {
	if s.cfg.RabbitMq == nil {
		panic("请在配置文件中配置 rabbitmq 的配置")
	}
	s.RabbitMQ = &xrabbitmq.AMQP{Conf: s.cfg.RabbitMq}
	s.RabbitMQ.PullLimit = s.cfg.Email.BatchNumber    // 每次最多 拉多少条
	s.RabbitMQ.ConsumerLimit = s.cfg.Email.Consumer   // 每次最多 多少个消费者
	s.RabbitMQ.Exchange = constant.RABBIT_MQ_EXCHANGE // 交换机名
	s.RabbitMQ.Queue = constant.RABBIT_MQ_QUEUE       // 队列名
	s.RabbitMQ.Start()
	s.RabbitMQ.QueueDeclare()
	s.RabbitMQ.BindRoutingKey(constant.RABBIT_MQ_ROUTIE_KEY) // 初始化约定要绑定的 routing_key
}

func (q *QueueEngineRabbitMq) Close(s *Service) {
	_ = s.RabbitMQ.Close()
}

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

func (q *QueueEngineKafka) Start(s *Service) {
	if s.cfg.KafkaTopic == nil {
		panic("请在配置文件中配置 kafkaTopic")
	}
	if s.cfg.Kafka == nil {
		panic("请在配置文件中配置 kafka")
	}
	if s.serviceType == constant.SERVICE_TYPE_API || s.serviceType == constant.SERVICE_TYPE_ALL {
		s.KafkaEmailProducer = xkafka.NewProducer(s.cfg.Kafka)
	}
	if s.serviceType == constant.SERVICE_TYPE_JOB || s.serviceType == constant.SERVICE_TYPE_ALL {
		s.KafkaEmailConsumer = xkafka.NewConsumer(s.cfg.Kafka, s.cfg.KafkaTopic.TopicList, s.cfg.KafkaTopic.Group)
	}
}

func (q *QueueEngineKafka) Close(s *Service) {
	if s.KafkaEmailConsumer != nil {
		_ = s.KafkaEmailConsumer.Consumer.Close()
	}
	if s.KafkaEmailProducer != nil {
		_ = s.KafkaEmailProducer.Close()
	}
}

func (q *QueueEngineKafka) Push(s *Service, payload []byte) (err error) {
	err = s.KafkaEmailProducer.SendMsgAsync(s.cfg.KafkaTopic.TopicList[0], payload)
	return err
}

func (q *QueueEngineKafka) Pull(s *Service, callback func([]byte) error) error {
	registerHandler := func(session *xkafka.ConsumerSession, msgs <-chan *sarama.ConsumerMessage) (errKafka error) {
		fun := xkafka.IteratorBatchFetch(session, msgs, s.cfg.Email.BatchNumber, constant.KAFKA_BATCH_WAIT_MAX_SECOND)
		for {
			kafkaMessages, ok := fun()
			if !ok {
				return
			}
			eg := &errgroup.Group{}
			eg.GOMAXPROCS(s.cfg.Email.Consumer)
			for _, business := range kafkaMessages {
				tmp := business
				eg.Go(func(context.Context) error {
					err := callback(tmp.Value)
					if err != nil {
						xlog.Errorf("Kafka Callback Value(%v) Group(%+v) Topic(%v) Err(%v)",
							string(tmp.Value), s.cfg.KafkaTopic.Group,
							strings.Join(s.cfg.KafkaTopic.TopicList, ","), err)
					} else {
						xlog.Infof("Kafka Callback Value(%v) Group(%+v) Topic(%v) Success",
							string(tmp.Value), s.cfg.KafkaTopic.Group,
							strings.Join(s.cfg.KafkaTopic.TopicList, ","))
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
