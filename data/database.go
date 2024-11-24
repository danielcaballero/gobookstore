package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDB establishes a connection to the MongoDB database
func ConnectDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return fmt.Errorf("failed to create MongoDB client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	DB = client.Database("bookstore")
	fmt.Println("Connected to MongoDB")
	return nil
}

// DisconnectDB closes the database connection
func DisconnectDB() {
	if DB != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = DB.Client().Disconnect(ctx)
	}
}
