package infra

import (
	"github.com/Postech-fiap-soat/ms-payment/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel(cfg *config.Config) (*amqp.Channel, error) {
	conn, err := amqp.Dial(cfg.RabbitDialStr)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}
