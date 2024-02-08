/*
 * @File: models.movie.go
 * @Description: Defines Movie information will be returned to the clients
 * 
 */
package models

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type Challenge struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Status 		string
}

type AddChallenge struct {
	Name        string `json:"name" example:"Movie Name"`
	Status 		string
}
