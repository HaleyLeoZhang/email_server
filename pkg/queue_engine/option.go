package queue_engine

// ----------------------------------------------------------------------
// 初始化包
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"github.com/HaleyLeoZhang/email_server/conf"
)

/**
 * 简单工厂
 *
 * @return queue_engine.Queue
 */
func GetEmailQueue() Queue {
	switch conf.Conf.Email.Driver {
	case "amqp":
		return &RabbitMq{}
	case "kafka":
		return &Kafka{}
	default:
		panic("驱动配置错误")
	}
}
