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
		_ = NewJsendPresenter(w, http.StatusBadRequest, err.Error(), nil)
		return err
	}
	var paymentDto domain.CreatePaymentInputDTO
	err = json.Unmarshal(body, &paymentDto)
	if err != nil {
		_ = NewJsendPresenter(w, http.StatusBadRequest, err.Error(), nil)
		return err
	}
	err = h.usecase.CreatePayment(r.Context(), paymentDto)
	if err != nil {
		_ = NewJsendPresenter(w, http.StatusInternalServerError, err.Error(), nil)
		return err
	}
	return NewJsendPresenter(w, http.StatusOK, "Payment approved successfully", nil)
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
