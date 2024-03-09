package auth_middleware

import (
	"context"
	"fmt"
	"net/http"
	"todo_app/src/config"
	user_models "todo_app/src/user/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Collection = config.ConnectDB().Collection("user")

func VerifyToken(tokenString string) (*user_models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &user_models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SecretKey"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*user_models.Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Geçersiz token")
	}

	return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "Login information was not provided."})
			c.Abort()
			return
		}

		claims, err := VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		filter := bson.D{{Key: "username", Value: claims.Username}}

		var user user_models.UserDetail

		err = db.FindOne(context.TODO(), filter).Decode(&user)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"detail": "User not found!"})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
