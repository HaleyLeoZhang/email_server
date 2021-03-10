package queue_engine

import "github.com/HaleyLeoZhang/email_server/service"

// ----------------------------------------------------------------------
// 接口限定
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Queue interface {
	Push(s *service.Service, payload []byte) error
	Pull(s *service.Service, callback func([]byte) error) error
}
