package core

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strings"
	utils "user_api/utils"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	input := utils.ReadBody(c.Request.Body)
	var data User
	var isUser = true
	json.Unmarshal([]byte(input), &data)
	token := utils.JWTAuthService().GenerateToken(data.Email, isUser)
	c.JSON(http.StatusOK, gin.H{"data": token})
}

func VerifyToken(c *gin.Context) {
	authHeader := strings.Split(c.GetHeader("Authorization"), " ")
	tokenString := authHeader[1]
	println(tokenString)
	token, err := utils.JWTAuthService().ValidateToken(tokenString)
	if token.Valid {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	} else {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "INVALID TOKEN"})
	}
}

// jwt
//auth-jwt

// go get github.com/dgrijalva/jwt-go
// go mod vendor ต้องรันหลังจากลงอะไรใหม่
