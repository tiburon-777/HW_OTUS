package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func (r *Rabbit) Publish(b []byte) error {
	ch, err := r.Connection.Channel()
	defer func() {
		if err := ch.Close(); err != nil {
			log.Println("can't close AMQP connection")
		}
	}()
	if err != nil {
		return fmt.Errorf("can't get channel from AMQP connection: %w", err)
	}
	err = ch.Publish(
		r.Exchange, // exchange
		r.Key,      // routing key
		true,       // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json; charset=utf-8",
			Body:         b,
		})
	if err != nil {
		return fmt.Errorf("can't publish message into RabbitMQ: %w", err)
	}
	return nil
}
