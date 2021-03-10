package service

import (
	"github.com/HaleyLeoZhang/email_server/conf"
	"github.com/HaleyLeoZhang/email_server/dao"
	"github.com/HaleyLeoZhang/email_server/http/email"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/HaleyLeoZhang/go-component/driver/xrabbitmq"
)

type Service struct {
	DB *dao.Dao
	RabbitMQ *xrabbitmq.AMQP
}

func New(cfg *conf.Config) *Service {
	s := &Service{}
	if cfg.DB != nil {
		s.DB = dao.New(cfg.DB)
	}
	if cfg.RabbitMq != nil {
		s.DB = dao.New(cfg.DB)
	}
	return s
}

func (s *Service) Start()  {
	go s.DoPull()
}

// Close close the resource.
func (s *Service) Close() {
	// 各种消费者
	// - 暂无
	// 各种数据库
	// - 平滑关闭，建议数据库相关的关闭放到最后
	if s.CacheDao != nil {
		s.CacheDao.Close()
	}
	if s.CurlAvatarDao != nil {
		s.CurlAvatarDao.Close()
	}
	xlog.Info("Close.commonService.Done")
}
