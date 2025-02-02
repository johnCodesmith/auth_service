package main

import (
	"auth-service/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Создаем новый роутер
	router := gin.Default()

	// Регистрируем маршруты и их обработчики
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	// Запускаем сервер на порту 8080
	log.Fatal(router.Run(":8080"))
}
