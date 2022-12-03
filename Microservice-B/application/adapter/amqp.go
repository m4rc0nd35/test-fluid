package adapter

import amqp "github.com/rabbitmq/amqp091-go"

type Channel struct {
	amqp.Channel
}

type Amqp interface {
	SendToQueu(queue string, txt string) error
	ConsumerQueue(queue string, prefetch int, callback func([]byte) bool) error
}
