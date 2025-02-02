package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4" // Обновлённая библиотека для работы с JWT
)

// 🔑 Секретный ключ для подписи токенов (его лучше хранить в ENV-переменных!)
var jwtSecret = []byte("my_secret_key")

// GenerateJWT создаёт JWT-токен для указанного пользователя
func GenerateJWT(username string) (string, error) {
	// 🏷 Определяем payload токена (данные, которые он несёт)
	claims := jwt.MapClaims{
		"username": username,                               // 👤 Записываем имя пользователя в токен
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // ⏳ Указываем срок действия токена (24 часа)
		"iat":      time.Now().Unix(),                     // ⏱ Добавляем время выпуска токена
	}

	// 🔐 Создаём новый токен с алгоритмом HMAC-SHA256 и нашими claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 🖋 Подписываем токен с использованием секретного ключа
	return token.SignedString(jwtSecret)
}
