package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GreetingEnglish(c *gin.Context) {
	name := c.GetHeader("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name header is required"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s!", name)})
}

func GreetingFrench(c *gin.Context) {
	name := c.GetHeader("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name header is required"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Bonjour, %s!", name)})
}

func GreetingSpanish(c *gin.Context) {
	name := c.GetHeader("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name header is required"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("¡Hola, %s!", name)})
}
