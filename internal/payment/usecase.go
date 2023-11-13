package payment

import (
	"context"
	"github.com/Postech-fiap-soat/ms-payment/internal/domain"
)

type Usecase struct {
	repository          domain.Repository
	service             domain.Service
	prodQueueRepository domain.ProdQueueRepository
}

func NewUseCase(repository domain.Repository, prodQueueRepository domain.ProdQueueRepository, service domain.Service) *Usecase {
	return &Usecase{repository: repository, service: service, prodQueueRepository: prodQueueRepository}
}

func (p *Usecase) CreatePayment(ctx context.Context, paymentDto domain.CreatePaymentInputDTO) error {
	payment := domain.NewPayment(paymentDto)
	err := p.service.ApplyAPIPayment(payment)
	if err != nil {
		return err
	}
	err = p.repository.CreatePayment(ctx, payment)
	if err != nil {
		return err
	}
	err = p.prodQueueRepository.PublishPayment(payment)
	if err != nil {
		return err
	}
	return nil
}
