package service

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Connection  *amqp.Connection
	IsConnected bool
}

func NewConnectAMQP(master, port, username, password, vhost string) (*RabbitMQ, error) {
	rabbitMQ := RabbitMQ{}
	var err error
	cfg := amqp.Config{
		Properties: amqp.Table{
			"connection_name": fmt.Sprint("DATA_WORKER-DATACENTER-MANAGER"),
		},
	}

	/* Connect AMQP */
	rabbitMQ.Connection, err = amqp.DialConfig(fmt.Sprint("amqp://", username, ":", password, "@", master, ":", port, "/", vhost), cfg)
	if err != nil {
		rabbitMQ.IsConnected = false
		return nil, err
	}

	rabbitMQ.IsConnected = true

	return &rabbitMQ, nil
}

func (c *RabbitMQ) SendToQueu(queue string, txt string) error {
	/* Check connection AMQP */
	if !c.IsConnected || c.Connection.IsClosed() {
		c.IsConnected = false
		time.Sleep(time.Second * 1)
		return fmt.Errorf("Connection is closed")
	}

	channel, err := c.Connection.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	q, err := channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		// channel.Close()
		return fmt.Errorf("QueueDeclare: %s", err)
	}

	err = channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(txt),
		},
	)
	if err != nil {
		// channel.Close()
		return fmt.Errorf("Publish: %s", err)
	}

	channel.Close()
	return nil
}

/*
(queue name string, callback name func(body receiver string, channel receiver *amqp.Channel, delivery tag uint64))
*/
func (c *RabbitMQ) ConsumerQueue(queue string, prefetch int, callback func(string, *amqp.Channel, uint64)) error {
	if !c.IsConnected {
		return fmt.Errorf("Connection is closed")
	}

	channel, err := c.Connection.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	channel.Qos(
		prefetch, // prefetch count
		0,        // prefetch size
		false,    // global
	)

	queueDeclared, err := channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	msgs, err := channel.Consume(
		queueDeclared.Name, // queue
		queueDeclared.Name, // consumer
		false,              // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)

	go func() {
		for d := range msgs {
			callback(string(d.Body), channel, d.DeliveryTag)
		}
	}()

	return nil
}
