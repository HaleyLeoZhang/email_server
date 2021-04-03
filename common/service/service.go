package service

import (
	"github.com/HaleyLeoZhang/email_server/common/constant"
	"github.com/HaleyLeoZhang/email_server/common/dao"
	"github.com/HaleyLeoZhang/email_server/common/model/bo"
	"github.com/HaleyLeoZhang/go-component/driver/xkafka"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/driver/xrabbitmq"
)

type Service struct {
	serviceType int // ALL、API、JOB
	cfg         *bo.Config
	DB          *dao.Dao
	// 消息驱动配置
	RabbitMQ           *xrabbitmq.AMQP
	KafkaEmailConsumer *xkafka.Consumer
	KafkaEmailProducer *xkafka.Producer
	// 邮件服务
	instanceEmailQueue Queue // 简单工厂获取的队列实例
}

func New(cfg *bo.Config, serviceType int) *Service {
	xlog.Init(cfg.Log)
	s := &Service{
		cfg:         cfg,
		serviceType: serviceType,
	}
	if cfg.DB != nil {
		s.DB = dao.New(cfg.DB)
	}
	s.instanceEmailQueue = s.GetEmailQueue()
	return s
}

func (s *Service) Start() {
	switch true {
	case s.serviceType == constant.SERVICE_TYPE_JOB || s.serviceType == constant.SERVICE_TYPE_ALL:
		go s.DoMessagePull()
	}
}

func (s *Service) Close() {
	// 各种消费者
	// - 暂无
	s.instanceEmailQueue.Close(s)
	// 各种数据库
	// - 平滑关闭，建议数据库相关的关闭放到最后
	if s.DB != nil {
		s.DB.Close()
	}
	xlog.Info("Close.Service.Done")
}
