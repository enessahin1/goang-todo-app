package controllers

import (
	"context"
	"net/http"

	"todo_app/src/todo/models"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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
