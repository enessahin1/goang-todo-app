package controllers

import (
	"todo_app/src/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Collection = config.ConnectDB().Collection("todos")
