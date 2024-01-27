package main

import (
	"context"
	"github.com/Postech-fiap-soat/ms-payment/internal/config"
	"github.com/Postech-fiap-soat/ms-payment/internal/infra"
	"github.com/Postech-fiap-soat/ms-payment/internal/payment"
	"github.com/uptrace/bunrouter"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(".")
	clientDB, err := infra.GetDatabaseConnection(ctx, cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer func() {
		if err = clientDB.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	queueCh, err := infra.OpenChannel(cfg)
	repository := payment.NewRepository(clientDB)
	prodQueueRepository := payment.NewProdQueueRepository(cfg, queueCh)
	service := payment.NewService()
	usecase := payment.NewUseCase(repository, prodQueueRepository, service)
	handler := payment.NewHandler(usecase)
	router := bunrouter.New()
	router.POST("/payment", handler.CreatePayment)
	log.Fatalf(http.ListenAndServe(":8001", router).Error())
}
