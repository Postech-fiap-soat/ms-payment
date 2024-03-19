package payment

import (
	"context"
	"fmt"
	"github.com/Postech-fiap-soat/ms-payment/internal/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db         *mongo.Client
	collection *mongo.Collection
}

func NewRepository(db *mongo.Client) domain.Repository {
	collection := db.Database("payment_historic").Collection("payments")
	return &Repository{db: db, collection: collection}
}

func (p *Repository) CreatePayment(ctx context.Context, payment *domain.Payment) error {
	document := bson.D{
		{"id", uuid.New().ID()},
		{"total_price", payment.TotalPrice},
		{"status", payment.Status},
		{"order", payment.Order},
		{"client_data", payment.Client},
	}
	result, err := p.collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	payment.ID = fmt.Sprintf("%v", result.InsertedID)
	return nil
}
