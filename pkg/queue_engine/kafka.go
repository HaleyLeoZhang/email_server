package queue_engine

import (
	"github.com/HaleyLeoZhang/email_server/service"
)

// ----------------------------------------------------------------------
// Kafka
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
// 包源 https://github.com/optiopay/kafka
// ----------------------------------------------------------------------

type Kafka struct{}

func (k *Kafka) Push(s *service.Service, payload []byte) error {
	return nil
}

func (k *Kafka) Pull(s *service.Service, callback func([]byte) error) error {
	return nil
}
