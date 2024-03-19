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
	TotalPrice *float64
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
	Cart          CartDTO   `json:"cart"`
	Client        ClientDTO `json:"client"`
	Observation   string    `json:"observation"`
	TotalPrice    float64   `json:"totalPrice"`
	PaymentStatus int       `json:"payment_status"`
	OrderStatus   int       `json:"order_status"`
}

type ClientDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Email string `json:"email"`
}

type CartDTO struct {
	Id    int       `json:"id"`
	Items []ItemDTO `json:"items"`
}

type ItemDTO struct {
	Id          int     `json:"id"`
	Count       int     `json:"count"`
	Product     Product `json:"product"`
	Observation string  `json:"observation"`
}

type Product struct {
	Id          int     `json:"id"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

func NewPayment(paymentDto CreatePaymentInputDTO) *Payment {
	p := Payment{
		TotalPrice: &paymentDto.TotalPrice,
		Status:     int64(paymentDto.OrderStatus),
		Order: &Order{
			ItemsTitle:     "",
			ItemsQuantity:  1,
			ItemsUnitPrice: paymentDto.TotalPrice,
		},
		Client: &Client{
			TypeIdentification:   "CPF",
			NumberIdentification: paymentDto.Client.Cpf,
			Name:                 paymentDto.Client.Name,
			Surname:              paymentDto.Client.Name,
			Email:                paymentDto.Client.Email,
		},
	}
	return &p
}

func (p *Payment) PaidSuccessfully() {
	p.Status = success
}
