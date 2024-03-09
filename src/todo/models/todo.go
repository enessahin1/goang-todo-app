package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoRequest struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedUser primitive.ObjectID `bson:"created_user"`
}

type TodoResponse struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedUser primitive.ObjectID `bson:"created_user"`
}
