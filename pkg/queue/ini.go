package queue

// ----------------------------------------------------------------------
// 初始化包
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

import (
	"email_server/pkg/e"
	"email_server/pkg/setting"
	"sync"
)

/**
 * 包内全局变量
 */
type oneInstacne struct {
	// 单例连接
	Conn interface{}
	// 因为多协程共用一个tcp链接,防止并发交错错写入
	// - 但一个连接能建立多个通道
	Lock sync.Mutex
	// 消费最大数量
	Pool chan int
}

var one oneInstacne

/**
 * 简单工厂
 *
 * @return queue.Queue
 */
func GetEmailQueue() Queue {
	one.Pool = make(chan int, setting.QueueSetting.CHANNEL_NUMBER)
	switch setting.QueueSetting.DRIVER {
	case "amqp":
		return &AMQP{
			Exchange: e.AMQP_MAIL_EXCHANGE,
			Queue:    e.AMQP_MAIL_QUEUE,
		}
	case "kafka":
		return &Kafka{}
	default:
		panic("驱动配置错误")
	}
}
