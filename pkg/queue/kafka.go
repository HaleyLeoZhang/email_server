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
	Payload []byte
}

func (q *Kafka) SetPayload(payload []byte) {
	q.Payload = payload
}

func (q *Kafka) Push() error {
	// q.Payload
	return nil
}

func (q *Kafka) Pull(callback func([]byte) error) error {
	// q.Payload
	return nil
}

func (a *Kafka) Close() error {
	return nil
}
