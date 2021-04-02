package service

import (
	"github.com/HaleyLeoZhang/email_server/conf"
	"github.com/HaleyLeoZhang/email_server/constant"
	"github.com/HaleyLeoZhang/email_server/dao"
	"github.com/HaleyLeoZhang/go-component/driver/xkafka"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/driver/xrabbitmq"
)

type Service struct {
	DB                 *dao.Dao
	RabbitMQ           *xrabbitmq.AMQP
	KafkaEmailConsumer *xkafka.Consumer
	KafkaEmailProducer *xkafka.Producer
}

func New() *Service {
	s := &Service{}
	if conf.Conf.DB != nil {
		s.DB = dao.New(conf.Conf.DB)
	}
	if conf.Conf.RabbitMq != nil {
		s.RabbitMQ = &xrabbitmq.AMQP{Conf: conf.Conf.RabbitMq}
	}
	if conf.Conf.Kafka != nil {
		if conf.Conf.KafkaTopic == nil {
			panic("请在配置文件中配置 kafkaTopic")
		}
		s.KafkaEmailConsumer = xkafka.NewConsumer(conf.Conf.Kafka, conf.Conf.KafkaTopic.TopicList, conf.Conf.KafkaTopic.Group)
		s.KafkaEmailProducer = xkafka.NewProducer(conf.Conf.Kafka)
	}
	return s
}

func (s *Service) Start() {
	// 初始化消费者
	switch true {
	case s.RabbitMQ != nil:
		s.RabbitMQ.PullLimit = conf.Conf.Email.BatchNumber  // 每次最多 拉多少条
		s.RabbitMQ.ConsumerLimit = conf.Conf.Email.Consumer // 每次最多 多少个消费者
		s.RabbitMQ.Exchange = constant.RABBIT_MQ_EXCHANGE   // 交换机名
		s.RabbitMQ.Queue = constant.RABBIT_MQ_QUEUE         // 队列名
		s.RabbitMQ.Start()
		s.RabbitMQ.QueueDeclare()
		s.RabbitMQ.BindRoutingKey(constant.RABBIT_MQ_ROUTIE_KEY) // 初始化约定要绑定的 routing_key
	}
	go s.DoMessagePull()
}

func (s *Service) Close() {
	// 各种消费者
	// - 暂无
	switch true {
	case s.RabbitMQ != nil:
		_ = s.RabbitMQ.Close()
	case s.KafkaEmailConsumer != nil:
		_ = s.KafkaEmailConsumer.Consumer.Close()
		_ = s.KafkaEmailProducer.Close()
	}
	// 各种数据库
	// - 平滑关闭，建议数据库相关的关闭放到最后
	if s.DB != nil {
		s.DB.Close()
	}
	xlog.Info("Close.commonService.Done")
}
