package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"social_network/config"
	"time"
)

type Conn struct {
	client *mongo.Client
	DB     *mongo.Database
}

func Connection(config *config.Config) (*Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.ConnectionUri())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		disconnectErr := client.Disconnect(ctx)
		if disconnectErr != nil {
			return nil, disconnectErr
		}
		return nil, err
	}

	return &Conn{client, client.Database(config.DBName())}, nil
}

func (c *Conn) Disconnect() error {
	if c == nil || c.client == nil {
		return errors.New("no connection to disconnect")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.client.Disconnect(ctx)
}
