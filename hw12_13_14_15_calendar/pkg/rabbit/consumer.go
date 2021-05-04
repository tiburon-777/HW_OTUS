package rabbit

import (
	"context"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RMQConnection interface {
	Channel() (*amqp.Channel, error)
}

type Message struct {
	Ctx  context.Context
	Data []byte
}

func (r *Rabbit) Consume(ctx context.Context, queue string) (<-chan Message, error) {
	messages := make(chan Message)
	ch, err := r.Connection.Channel()
	if err != nil {
		return nil, fmt.Errorf("can't get channel from AMQP connection: %w", err)
	}
	deliveries, err := ch.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("start consuming: %w", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case del := <-deliveries:
				if err := del.Ack(false); err != nil {
					log.Println(err)
				}

				msg := Message{
					Ctx:  context.TODO(),
					Data: del.Body,
				}

				select {
				case <-ctx.Done():
					return
				case messages <- msg:
				}
			}
		}
	}()

	return messages, nil
}
