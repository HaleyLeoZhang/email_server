package queue

// ----------------------------------------------------------------------
// Kafka
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
// 包源 https://github.com/optiopay/kafka
// ----------------------------------------------------------------------

type Kafka struct {
	Payload  string
	Callback func(string)
}

func (q *Kafka) Push() error {
	// q.Payload
	return nil
}

func (q *Kafka) Pull(callback func(string) error) error {
	// q.Payload
	return nil
}
