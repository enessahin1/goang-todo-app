package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
