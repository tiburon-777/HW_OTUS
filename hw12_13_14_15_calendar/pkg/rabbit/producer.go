package rabbit

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func (r *Rabbit) Publish(body string) error {
	b, err := json.Marshal([]byte(body))
	if err != nil {
		return err
	}
	err = r.Channel.Publish(
		r.Exchange, // exchange
		r.Key,      // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json; charset=utf-8",
			Body:        b,
		})
	if err != nil {
		return err
	}
	return nil
}
