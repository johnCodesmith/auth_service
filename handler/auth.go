package handler
import (
	"auth-service/model"   // Пакет для работы с моделями данных (например, структура User)
	"auth-service/service" // Пакет для логики аутентификации (например, сервисы регистрации и логина)
	"github.com/gin-gonic/gin" // Используем Gin для создания веб-сервиса
	"net/http"  // Стандартная библиотека для HTTP-статусов и работы с HTTP-запросами
)

// Инициализация сервиса аутентификации
// Это объект, который инкапсулирует логику регистрации и логина пользователей.
var authService = service.NewAuthService()

// Register — функция для регистрации пользователя.
func Register(c *gin.Context) {
	var user model.User // Структура для пользователя, которую будем заполнять из запроса

	// Проверяем, что JSON-данные из запроса корректно парсятся в структуру User
	// Если есть ошибка в данных (например, не хватает обязательных полей), возвращаем ошибку 400
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращаем ошибку с подробным сообщением
		return
	}

	// Вызываем сервис для регистрации пользователя
	err := authService.Register(user)

	// Если произошла ошибка при регистрации (например, такой пользователь уже существует), возвращаем ошибку 500
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Если регистрация успешна, отправляем ответ с сообщением об успешной регистрации
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login — функция для логина пользователя.
func Login(c *gin.Context) {
	var user model.User // Структура для пользователя, которую будем заполнять из запроса

	// Проверяем, что JSON-данные из запроса корректно парсятся в структуру User
	// Если есть ошибка в данных, возвращаем ошибку 400
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращаем ошибку с подробным сообщением
		return
	}

	// Вызываем сервис для выполнения логина и получения JWT-токена
	token, err := authService.Login(user)

	// Если возникла ошибка (например, неверный логин или пароль), возвращаем ошибку 401 (Unauthorized)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Если логин успешен, возвращаем токен
	c.JSON(http.StatusOK, gin.H{"token": token})
}
