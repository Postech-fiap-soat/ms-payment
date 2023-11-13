package payment

import (
	"encoding/json"
	"github.com/Postech-fiap-soat/ms-payment/internal/domain"
	"github.com/uptrace/bunrouter"
	"io"
	"net/http"
)

type Handler struct {
	usecase domain.Usecase
}

func NewHandler(usecase domain.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) CreatePayment(w http.ResponseWriter, r bunrouter.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	var paymentDto domain.CreatePaymentInputDTO
	err = json.Unmarshal(body, &paymentDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	err = h.usecase.CreatePayment(r.Context(), paymentDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.Write(body)
	return nil
}
