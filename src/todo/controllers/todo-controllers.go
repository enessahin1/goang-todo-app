package controllers

import (
	"context"
	"net/http"

	"todo_app/src/config"
	"todo_app/src/todo/models"
	user_models "todo_app/src/user/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Collection = config.ConnectDB().Collection("todos")

func CreateTodo(c *gin.Context) {
	var data models.TodoRequest
	var created_user user_models.UserDetail

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created_user, ok := c.Keys["user"].(user_models.UserDetail)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user"})
		return
	}

	todo := models.TodoRequest{
		Name:        data.Name,
		Description: data.Description,
		CreatedUser: created_user.ID,
	}

	created_data, err := db.InsertOne(context.TODO(), &todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response models.TodoResponse
	response.ID = created_data.InsertedID.(primitive.ObjectID)
	response.Name = todo.Name
	response.Description = todo.Description

	c.JSON(http.StatusCreated, response)
}

func GetAllTodos(c *gin.Context) {
	var result []models.TodoResponse

	cursor, err := db.Find(context.TODO(), bson.M{})

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, gin.H{"body": "Data yok"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for cursor.Next(context.Background()) {
		var elem models.TodoResponse
		err := cursor.Decode(&elem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result = append(result, elem)
	}

	c.JSON(http.StatusOK, result)
}

func RetriveTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var result models.TodoResponse

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"detail": "Not Found"})
		return
	}

	err = db.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"detail": "Not Found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateTodo(c *gin.Context) {
	var data models.TodoRequest
	id, _ := primitive.ObjectIDFromHex(cast.ToString(c.Params.ByName("id")))
	filter := bson.D{{"_id", id}}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request := models.TodoRequest{Name: data.Name, Description: data.Description}
	_, err := db.UpdateOne(context.TODO(), filter, bson.D{{"$set", request}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := models.TodoResponse{
		ID:          id,
		Name:        data.Name,
		Description: data.Description,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteTodo(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(cast.ToString(c.Params.ByName("id")))
	filter := bson.D{{"_id", id}}

	_, err := db.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
	})
}
