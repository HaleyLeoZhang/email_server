package queue

// ----------------------------------------------------------------------
// 接口限定
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"email_server/pkg/e"
)

type Queue interface {
	Push() error
	Pull(callback func(string) error) error
}

func GetEmailQueue() *AMQP {
	return &AMQP{
		Exchange: e.AMQP_MAIL_EXCHANGE,
		Queue:    e.AMQP_MAIL_QUEUE,
	}
}
