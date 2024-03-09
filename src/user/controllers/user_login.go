package user_login

import (
	"context"
	"net/http"
	"time"
	"todo_app/src/config"
	user_models "todo_app/src/user/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var db *mongo.Collection = config.ConnectDB().Collection("user")

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(user user_models.User) (string, error) {
	claims := user_models.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("SecretKey"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

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

	user_match := CheckPasswordHash(user_model.Password, filtered_user.Password)

	if user_match == false {
		c.JSON(http.StatusNotFound, gin.H{"detail": "Username or password incorrect!"})
		return
	}

	generated_token, err := GenerateToken(filtered_user)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Token could not generated!", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": generated_token})

}

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
