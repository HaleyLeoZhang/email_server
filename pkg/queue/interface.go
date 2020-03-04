package queue

// ----------------------------------------------------------------------
// 接口限定
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------

type Queue interface {
	Push() error
	Pull(callback func(string) error) error
	SetPayload([]byte)
}
