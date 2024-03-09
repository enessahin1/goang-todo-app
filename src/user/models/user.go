package user_models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	IsActive    bool   `json:"is_active"`
	IsSuperuser bool   `json:"is_superuser"`
	CreatedDate string `json:"created_date"`
	LastLogin   string `json:"last_login"`
}

type UserDetail struct {
	ID          primitive.ObjectID `bson:"_id"`
	Username    string             `json:"username"`
	Password    string             `json:"password"`
	IsActive    bool               `json:"is_active"`
	IsSuperuser bool               `json:"is_superuser"`
	CreatedDate string             `json:"created_date"`
	LastLogin   string             `json:"last_login"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
