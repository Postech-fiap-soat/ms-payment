package payment

import (
	"errors"
	"github.com/Postech-fiap-soat/ms-payment/internal/config"
	"github.com/Postech-fiap-soat/ms-payment/internal/domain"
	"github.com/eduardo-mior/mercadopago-sdk-go"
)

type Service struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) domain.Service {
	return &Service{cfg: cfg}
}

func (s *Service) ApplyAPIPayment(payment *domain.Payment) (*domain.Payment, error) {
	_, mercadopagoErr, err := mercadopago.CreatePayment(mercadopago.PaymentRequest{
		ExternalReference: payment.ID,
		Items: []mercadopago.Item{
			{
				Title:     payment.Order.ItemsTitle,
				Quantity:  payment.Order.ItemsQuantity,
				UnitPrice: payment.Order.ItemsUnitPrice,
			},
		},
		Payer: mercadopago.Payer{
			Identification: mercadopago.PayerIdentification{
				Type:   payment.Client.TypeIdentification,
				Number: payment.Client.NumberIdentification,
			},
			Name:    payment.Client.Name,
			Surname: payment.Client.Surname,
			Email:   payment.Client.Email,
		},
		NotificationURL: s.cfg.WebhookNotification,
	}, s.cfg.MercadoPagoAccessToken)
	if err != nil {
		return nil, err
	} else if mercadopagoErr != nil {
		return nil, errors.New(mercadopagoErr.Error)
	}
	payment.PaidSuccessfully()
	return payment, nil
}
