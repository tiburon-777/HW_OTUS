package rabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Rabbit struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Exchange   string
	Key        string
	Queue      string
}

type Config struct {
	Login    string
	Pass     string
	Address  string
	Port     string
	Exchange string
	Queue    string
	Key      string
}

func New(conf Config) (*Rabbit, error) {
	conn, err := amqp.Dial("amqp://" + conf.Login + ":" + conf.Pass + "@" + conf.Address + ":" + conf.Port + "/")
	if err != nil {
		return nil, fmt.Errorf("can't dial RabbitMQ over AMQP: %w", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("can't get channel from AMQP connection: %w", err)
	}
	err = ch.ExchangeDeclare(
		conf.Exchange, // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("can't (re)declare exchange in RabbitMQ: %w", err)
	}
	q, err := ch.QueueDeclare(
		conf.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("can't (re)create queue in RabbitMQ: %w", err)
	}
	err = ch.QueueBind(q.Name, conf.Key, conf.Exchange, false, nil)
	if err != nil {
		return nil, fmt.Errorf("can't bind Queue on Exchange in RabbitMQ: %w", err)
	}
	return &Rabbit{Connection: conn, Channel: ch}, nil
}

func Attach(conf Config) (*Rabbit, error) {
	conn, err := amqp.Dial("amqp://" + conf.Login + ":" + conf.Pass + "@" + conf.Address + ":" + conf.Port + "/")
	if err != nil {
		return nil, fmt.Errorf("can't dial RabbitMQ over AMQP: %w", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("can't get channel from AMQP connection: %w", err)
	}
	return &Rabbit{Connection: conn, Channel: ch, Exchange: conf.Exchange, Queue: conf.Queue, Key: conf.Key}, nil
}

func (r *Rabbit) Close() error {
	if err := r.Channel.Close(); err != nil {
		return fmt.Errorf("can't close connection channel: %w", err)
	}
	if err := r.Connection.Close(); err != nil {
		return fmt.Errorf("can't close connection: %w", err)
	}
	return nil
}
