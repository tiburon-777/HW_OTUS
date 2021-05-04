package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Rabbit struct {
	Connection *amqp.Connection
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
	defer func() {
		if err := ch.Close(); err != nil {
			log.Println("can't close AMQP connection")
		}
	}()
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
	return &Rabbit{Connection: conn, Exchange: conf.Exchange, Key: conf.Key, Queue: conf.Queue}, nil
}

func Attach(conf Config) (*Rabbit, error) {
	conn, err := amqp.Dial("amqp://" + conf.Login + ":" + conf.Pass + "@" + conf.Address + ":" + conf.Port + "/")
	if err != nil {
		return nil, fmt.Errorf("can't dial RabbitMQ over AMQP: %w", err)
	}
	return &Rabbit{Connection: conn, Exchange: conf.Exchange, Queue: conf.Queue, Key: conf.Key}, nil
}

func (r *Rabbit) Close() error {
	if err := r.Connection.Close(); err != nil {
		return fmt.Errorf("can't close connection: %w", err)
	}
	return nil
}
