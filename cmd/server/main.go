package main

import (
	"context"
	"github.com/Postech-fiap-soat/ms-payment/internal/infra"
	"github.com/Postech-fiap-soat/ms-payment/internal/payment"
	"github.com/uptrace/bunrouter"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	clientDB, err := infra.GetDatabaseConnection(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	repository := payment.NewRepository(clientDB)
	prodQueueRepository := payment.NewProdQueueRepository()
	service := payment.NewService()
	usecase := payment.NewUseCase(repository, prodQueueRepository, service)
	handler := payment.NewHandler(usecase)
	router := bunrouter.New()
	router.POST("/payment", handler.CreatePayment)
	log.Fatalf(http.ListenAndServe(":8001", router).Error())
}
