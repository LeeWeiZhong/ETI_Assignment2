/*
 * @File: databases.mongodb.go
 * @Description: Handles MongoDB connections
 * 
 */
package databases

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	Client       *mongo.Client
	Databasename string
}

// Init initializes mongo database
func (db *MongoDB) Init() error {
	db.Databasename = "go"

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Can't connect to MongoDB: %v", err)
		return err
	}

	// Ping MongoDB to ensure connection is established
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Can't ping MongoDB: %v", err)
		client.Disconnect(context.Background())
		return err
	}

	db.Client = client

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.Client != nil {
		db.Client.Disconnect(context.Background())
	}
}
