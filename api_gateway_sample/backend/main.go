package main

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/greeting/en", handlers.GreetingEnglish)
	r.POST("/greeting/fr", handlers.GreetingFrench)
	r.POST("/greeting/es", handlers.GreetingSpanish)

	r.Run(":8080")
}
