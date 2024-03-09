package user_controller

import (
	"context"
	"net/http"
	"time"
	user_models "todo_app/src/user/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user_model user_models.User
	if err := c.ShouldBindJSON(&user_model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid data!"})
	}

	hassed_pass, err := bcrypt.GenerateFromPassword([]byte(user_model.Password), 14)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Password is not valid!"})
	}

	user_data := user_models.User{
		Username:    user_model.Username,
		Password:    string(hassed_pass),
		CreatedDate: time.Now().UTC().String(),
		IsSuperuser: false,
		IsActive:    true,
		LastLogin:   "",
	}

	_, err = db.InsertOne(context.TODO(), &user_data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user_data)
}
