package controllers

import (
	"context"
	"net/http"

	"todo_app/src/todo/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
