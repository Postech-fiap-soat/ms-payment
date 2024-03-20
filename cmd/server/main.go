package main

import (
	"context"
	"github.com/Postech-fiap-soat/ms-payment/internal/config"
	"github.com/Postech-fiap-soat/ms-payment/internal/infra"
	"github.com/Postech-fiap-soat/ms-payment/internal/payment"
	"log"
)

func main() {

	ctx := context.Background()
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	log.Println("Inicializando aplicação")
	LoadAPP(ctx, cfg)
}

func LoadAPP(ctx context.Context, cfg *config.Config) {
	clientDB, err := infra.GetDatabaseConnection(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = clientDB.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	queueCh, err := infra.OpenChannel(cfg)
	if err != nil {
		panic(err)
	}
	repository := payment.NewRepository(clientDB)
	prodQueueRepository := payment.NewProdQueueRepository(cfg, queueCh)
	service := payment.NewService(cfg)
	usecase := payment.NewUseCase(repository, prodQueueRepository, service)
	handler := payment.NewHandler(usecase)
	msgs, err := queueCh.Consume(
		"queue_pedidos",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		msg := <-msgs
		err = handler.CreatePayment(ctx, msg)
		if err != nil {
			log.Println("erro:", err)
		}
	}
}
