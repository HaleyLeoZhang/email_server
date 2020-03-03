package queue

// ----------------------------------------------------------------------
// 接口限定
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"email_server/pkg/e"
	"email_server/pkg/setting"
)

type Queue interface {
	Push() error
	Pull(callback func(string) error) error
	SetPayload([]byte)
}

/**
 * 简单工厂
 *
 * @return queue.Queue
 */
func GetEmailQueue() Queue {
	switch setting.QueueSetting.DRIVER {
	case "amqp":
		return &AMQP{
			Exchange: e.AMQP_MAIL_EXCHANGE,
			Queue:    e.AMQP_MAIL_QUEUE,
		}
	default:
		panic("驱动配置错误")
	}
}
