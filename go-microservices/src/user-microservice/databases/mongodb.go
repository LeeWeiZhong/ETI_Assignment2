package databases

import (
	"context"
	"log"
	"user-microservice/common"
	"user-microservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	Client       *mongo.Client
	DatabaseName string
}

// Init initializes mongo database
func (db *MongoDB) Init() error {
	db.DatabaseName = common.Config.MgDbName

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

	return db.initData()
}

// InitData initializes default data
func (db *MongoDB) initData() error {
	collection := db.Client.Database("go").Collection("users")
	
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatalf("Init: %v", err)
		return err
	}

	if count < 1 {
		// Create admin/admin account
		user := models.User{ID: primitive.NewObjectID(), Name: "admin", Password: "admin"}
		_, err = collection.InsertOne(context.Background(), user)
	}

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.Client != nil {
		db.Client.Disconnect(context.Background())
	}
}
