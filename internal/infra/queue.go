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
	ch.ExchangeDeclare("ex_pedidos", "direct", true, false, false, false, nil)
	ch.QueueDeclare("queue_pedidos", true, false, false, false, nil)
	ch.QueueBind("queue_pedidos", "key_pedidos", "ex_pedidos", false, nil)
	ch.ExchangeDeclare("ex_producao", "direct", true, false, false, false, nil)
	ch.QueueDeclare("queue_producao", true, false, false, false, nil)
	ch.QueueBind("queue_producao", "soatkey", "ex_producao", false, nil)
	return ch, nil
}
