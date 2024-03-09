package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	user_models "todo_app/src/user/models"
	user_utils "todo_app/src/user/utils"
)

func Login(c *gin.Context) {
	var user_model user_models.User
	if err := c.ShouldBindJSON(&user_model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Format"})
		return
	}

	var filtered_user user_models.User

	filter := bson.D{{"username", user_model.Username}}

	err := db.FindOne(context.TODO(), filter).Decode(&filtered_user)

	if err == mongo.ErrNoDocuments || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"detail": "Username or password incorrect!"})
		return
	}

	if !user_utils.CheckPasswordHash(user_model.Password, filtered_user.Password) {
		c.JSON(http.StatusNotFound, gin.H{"detail": "Username or password incorrect!"})
		return
	}

	generated_token, err := user_utils.GenerateToken(filtered_user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Token could not be generated!", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": generated_token})
}
