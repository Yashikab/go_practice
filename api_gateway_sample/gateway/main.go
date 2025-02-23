package main

import (
	"gateway/config"
	"gateway/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	router := handlers.NewRouter(cfg)
	r := gin.Default()

	r.POST("/greeting", router.RouteGreeting)

	r.Run(":8081")
}
