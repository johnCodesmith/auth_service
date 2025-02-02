package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4" 

var jwtSecret = []byte("my_secret_key")

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,                               
		"exp":      time.Now().Add(time.Hour * 24).Unix(), 
		"iat":      time.Now().Unix(),                     
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
