package handler
import (
	"auth-service/model"   
	"auth-service/service" 
	"github.com/gin-gonic/gin" 
	"net/http"  
)

var authService = service.NewAuthService()

func Register(c *gin.Context) {
	var user model.User 

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
		return
	}

	err := authService.Register(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user model.User 

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
		return
	}

	token, err := authService.Login(user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
