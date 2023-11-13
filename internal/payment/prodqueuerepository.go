package payment

import "github.com/Postech-fiap-soat/ms-payment/internal/domain"

type ProdQueueRepository struct {
}

func NewProdQueueRepository() domain.ProdQueueRepository {
	return &ProdQueueRepository{}
}

func (p *ProdQueueRepository) PublishPayment(payment *domain.Payment) error {
	return nil
}
