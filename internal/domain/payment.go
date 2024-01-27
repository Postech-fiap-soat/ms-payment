package domain

import (
	"context"
)

type Repository interface {
	CreatePayment(ctx context.Context, payment *Payment) error
}

type ProdQueueRepository interface {
	PublishPayment(ctx context.Context, payment *Payment) error
}

type Usecase interface {
	CreatePayment(ctx context.Context, paymentDto CreatePaymentInputDTO) error
}

type Service interface {
	ApplyAPIPayment(payment *Payment) (*Payment, error)
}

type Payment struct {
	ID         string
	OrderId    string
	TotalPrice *int64
	Status     int64
	Order      *Order  `json:"order"`
	Client     *Client `json:"client"`
}

type Client struct {
	TypeIdentification   string
	NumberIdentification string
	Name                 string
	Surname              string
	Email                string
}

type Order struct {
	ItemsTitle     string
	ItemsQuantity  float64
	ItemsUnitPrice float64
}

const (
	pending = 1
	success = 2
)

type CreatePaymentInputDTO struct {
	OrderId    string                      `json:"order_id"`
	TotalPrice *int64                      `json:"total_price"`
	Order      map[string]interface{}      `json:"order"`
	ClientData CreatePaymentClientInputDTO `json:"client"`
}

type CreatePaymentClientInputDTO struct {
	Cpf        string `json:"cpf"`
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
}

func NewPayment(paymentDto CreatePaymentInputDTO) *Payment {
	p := Payment{
		OrderId:    paymentDto.OrderId,
		TotalPrice: paymentDto.TotalPrice,
		Status:     pending,
		Order: &Order{
			ItemsTitle:     "",
			ItemsQuantity:  1,
			ItemsUnitPrice: float64(*paymentDto.TotalPrice),
		},
		Client: &Client{
			TypeIdentification:   "CPF",
			NumberIdentification: paymentDto.ClientData.Cpf,
			Name:                 paymentDto.ClientData.Name,
			Surname:              paymentDto.ClientData.SecondName,
			Email:                paymentDto.ClientData.Email,
		},
	}
	return &p
}

func (p *Payment) PaidSuccessfully() {
	p.Status = success
}
