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
	ID         string   `json:"id"`
	OrderId    int64    `json:"order_id"`
	TotalPrice *float64 `json:"total_price"`
	Status     int64    `json:"status"`
	Order      *Order   `json:"order"`
	Client     *Client  `json:"client"`
}

type Client struct {
	TypeIdentification   string `json:"type_identification"`
	NumberIdentification string `json:"number_identification"`
	Name                 string `json:"name"`
	Surname              string `json:"surname"`
	Email                string `json:"email"`
}

type Order struct {
	ItemsTitle     string  `json:"items_title"`
	ItemsQuantity  float64 `json:"items_quantity"`
	ItemsUnitPrice float64 `json:"items_unit_price"`
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
	Id    int64     `json:"id"`
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
		OrderId:    paymentDto.Cart.Id,
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
