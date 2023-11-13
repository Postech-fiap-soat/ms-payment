package domain

import "context"

type Repository interface {
	CreatePayment(ctx context.Context, payment *Payment) error
}

type ProdQueueRepository interface {
	PublishPayment(payment *Payment) error
}

type Usecase interface {
	CreatePayment(ctx context.Context, paymentDto CreatePaymentInputDTO) error
}

type Service interface {
	ApplyAPIPayment(payment *Payment) error
}

type Payment struct {
	ID         string
	TotalPrice *int64
	Status     int64
	Order      Order
	ClientData map[string]interface{} `json:"client_data"`
}

type Order struct {
	Items map[string]interface{}
}

const (
	pending = 1
)

type CreatePaymentInputDTO struct {
	ID         string                 `json:"id"`
	TotalPrice *int64                 `json:"total_price"`
	Order      map[string]interface{} `json:"order"`
	ClientData map[string]interface{} `json:"client_data"`
}

func NewPayment(paymentDto CreatePaymentInputDTO) *Payment {
	order := Order{}
	items, ok := paymentDto.Order["items"]
	if ok {
		order.Items = items.(map[string]interface{})
	}
	p := Payment{
		ID:         paymentDto.ID,
		TotalPrice: paymentDto.TotalPrice,
		Status:     pending,
		Order:      order,
		ClientData: paymentDto.ClientData,
	}
	return &p
}
