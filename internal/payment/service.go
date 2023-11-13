package payment

import "github.com/Postech-fiap-soat/ms-payment/internal/domain"

type Service struct {
}

func NewService() domain.Service {
	return &Service{}
}

func (s *Service) ApplyAPIPayment(payment *domain.Payment) error {
	return nil
}
