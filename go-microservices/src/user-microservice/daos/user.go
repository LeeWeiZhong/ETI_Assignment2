/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * 
 */
package daos

import (
	"context"
	//"user-microservice/common"
	"user-microservice/databases"
	"user-microservice/models"
	"user-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// User manages User CRUD
type User struct {
	utils *utils.Utils
}



func (u *User) GetAll() ([]models.User, error) {
	client := databases.Database.Client
	collection := client.Database("go").Collection("users")
	ctx := context.TODO()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []models.User
	err = cur.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}


// GetByID finds a User by its id
func (u *User) GetByID(id string) (models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	client := databases.Database.Client
	collection := client.Database("go").Collection("users")
	ctx := context.TODO()

	filter := bson.D{{"_id", objectID}}

	var user models.User
	err2 := collection.FindOne(ctx, filter).Decode(&user)
	if err2 != nil {
		if err2 == mongo.ErrNoDocuments {
			return models.User{}, err2
		}
		return models.User{}, err2
	}

	return user, nil
}

// DeleteByID deletes a User by its id
func (u *User) DeleteByID(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	client := databases.Database.Client
	collection := client.Database("go").Collection("users")
	ctx := context.TODO()

	filter := bson.D{{"_id", objectID}}

	_, err2 := collection.DeleteOne(ctx, filter)
	return err2
}

// Login authenticates a user based on username and password
func (u *User) Login( username, password string) (*User, error) {
	client := databases.Database.Client
	collection := client.Database("go").Collection("users")
	ctx := context.TODO()

	filter := bson.D{
		{"username", username},
		{"password", password},
	}

	var user User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

// Insert inserts a new User into the database
func (u *User) Insert(user models.User) error {
	client := databases.Database.Client
	collection := client.Database("go").Collection("users")
	ctx := context.TODO()

	_, err := collection.InsertOne(ctx, user)
	return err
}

// Delete deletes a User from the database
func (u *User) Delete(user models.User) error {
	client := databases.Database.Client
	collection := client.Database("go").Collection("users")
	ctx := context.TODO()

	filter := bson.D{{"_id", user.ID}}

	_, err := collection.DeleteOne(ctx, filter)
	return err
}

// Update updates an existing User in the database
func (u *User) Update(user models.User) error {
	client := databases.Database.Client
	collection := client.Database("go").Collection("users")
	ctx := context.TODO()

	filter := bson.D{{"_id", user.ID}}

	update := bson.D{
		{"$set", bson.D{
			{"Name", user.Name},
			{"Password", user.Password},
			// Add more fields as needed
		}},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}