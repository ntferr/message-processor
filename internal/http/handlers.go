package http

import (
	"message-processor/internal/controllers"
	"message-processor/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "It's all right",
	})
}

func CreateLogin(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controllers.CreateLogin(user)

	c.JSON(http.StatusOK, gin.H{"message": "User created!"})
}
