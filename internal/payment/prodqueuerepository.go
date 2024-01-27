package payment

import (
	"context"
	"encoding/json"
	"github.com/Postech-fiap-soat/ms-payment/internal/config"
	"github.com/Postech-fiap-soat/ms-payment/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ProdQueueRepository struct {
	queueCh *amqp.Channel
	cfg     *config.Config
}

func NewProdQueueRepository(cfg *config.Config, queueCh *amqp.Channel) domain.ProdQueueRepository {
	return &ProdQueueRepository{cfg: cfg, queueCh: queueCh}
}

func (p *ProdQueueRepository) PublishPayment(ctx context.Context, payment *domain.Payment) error {
	body, err := json.Marshal(payment)
	if err != nil {
		return err
	}
	err = p.queueCh.PublishWithContext(ctx, p.cfg.RabbitExchange, p.cfg.RabbitKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	if err != nil {
		return err
	}
	return nil
}
