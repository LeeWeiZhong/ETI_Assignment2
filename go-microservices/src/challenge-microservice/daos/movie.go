/*
 * @File: daos.movie.go
 * @Description: Implements Movie CRUD functions for MongoDB
 * 
 */
package daos

import (
	"context"	
	"challenge-microservice/databases"
	"challenge-microservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Challenge struct {
}

// COLLECTION of the database table
const (
	COLLECTION = "challenges"
)

// GetAll gets the list of Movie
func (m *Challenge) GetAll() ([]models.Challenge, error) {
	client := databases.Database.Client
	collection := client.Database("go").Collection("challenges")
	ctx := context.TODO()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var challenges []models.Challenge
	err = cur.All(ctx, &challenges)
	if err != nil {
		return nil, err
	}

	return challenges, nil
}

// GetByID finds a Movie by its id
func (m *Challenge) GetByID(id string) (models.Challenge, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Challenge{}, err
	}
	client := databases.Database.Client
	collection := client.Database("go").Collection("challenges")
	ctx := context.TODO()

	filter := bson.D{{"_id", objectID}}

	var challenge models.Challenge
	err2 := collection.FindOne(ctx, filter).Decode(&challenge)
	if err2 != nil {
		if err2 == mongo.ErrNoDocuments {
			return models.Challenge{}, err2
		}
		return models.Challenge{}, err2
	}

	return challenge, nil
}

// Insert adds a new Movie into database'
func (m *Challenge) Insert(challenge models.Challenge) error {
	client := databases.Database.Client
	collection := client.Database("go").Collection("challenges")
	ctx := context.TODO()

	_, err := collection.InsertOne(ctx, challenge)
	return err
}

// Delete remove an existing Movie
func (m *Challenge) Delete(challenge models.Challenge) error {
	client := databases.Database.Client
	collection := client.Database("go").Collection("challenges")
	ctx := context.TODO()

	filter := bson.D{{"_id", challenge.ID}}

	_, err := collection.DeleteOne(ctx, filter)
	return err
}

// Update modifies an existing Movie
func (m *Challenge) Update(challenge models.Challenge) error {
	client := databases.Database.Client
	collection := client.Database("go").Collection("challenges")
	ctx := context.TODO()

	filter := bson.D{{"_id", challenge.ID}}

	update := bson.D{
		{"$set", bson.D{
			{"Name", challenge.Name},
			{"Status", challenge.Status},
			// Add more fields as needed
		}},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
