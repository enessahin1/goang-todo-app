package controllers

import (
	"context"
	"net/http"
	"todo_app/src/todo/models"
	user_models "todo_app/src/user/models"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
