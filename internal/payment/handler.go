package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Postech-fiap-soat/ms-payment/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
)

type Handler struct {
	usecase domain.Usecase
}

func NewHandler(usecase domain.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) CreatePayment(ctx context.Context, msg amqp.Delivery) error {
	fmt.Println("foi")
	var paymentDto domain.CreatePaymentInputDTO
	err := json.Unmarshal(msg.Body, &paymentDto)
	if err != nil {
		return err
	}
	err = h.usecase.CreatePayment(ctx, paymentDto)
	if err != nil {
		return err
	}
	msg.Ack(true)
	return nil
}

type JsendPaymentApproved struct {
	Code    int
	Message string
	Data    interface{}
}

func NewJsendPresenter(w http.ResponseWriter, code int, message string, data interface{}) error {
	myjsend := JsendPaymentApproved{
		Code:    code,
		Message: message,
		Data:    data,
	}
	myJsendJson, err := json.Marshal(myjsend)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(myJsendJson)
	return nil
}
