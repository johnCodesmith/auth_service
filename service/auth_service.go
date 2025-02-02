package service

import (
	"auth-service/model" 
	"auth-service/utils" 
	"errors"
)

var userDB = make(map[string]string)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Register(user model.User) error {
	if _, exists := userDB[user.Username]; exists {
		return errors.New("user already exists")
	}

	userDB[user.Username] = user.Password
	return nil
}

func (a *AuthService) Login(user model.User) (string, error) {
	storedPassword, exists := userDB[user.Username]
	if !exists || storedPassword != user.Password {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
