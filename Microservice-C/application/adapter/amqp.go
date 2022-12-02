package adapter

type Amqp interface {
	SendToQueu(queue string, txt string) error
	ConsumerQueue(queue string, prefetch int, callback func([]byte) bool) error
}
