package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	mongoURL := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}

func CloseDB(client *mongo.Client) error {
	err := client.Disconnect(context.Background())
	if err != nil {
		return fmt.Errorf("error disconnecting from MongoDB: %v", err)
	}

	fmt.Println("Disconnected from MongoDB!")

	return nil
}
