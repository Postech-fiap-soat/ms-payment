package infra

import (
	"context"
	"github.com/Postech-fiap-soat/ms-payment/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetDatabaseConnection(ctx context.Context, cfg *config.Config) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.ConnStr))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
