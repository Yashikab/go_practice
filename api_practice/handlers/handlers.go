package handlers

import (
	"math/rand"
	"net/http"

	"api_practice/config"

	"github.com/gin-gonic/gin"
)

// レスポンス用の構造体
type InfoResponse struct {
	Version string `json:"version"`
}

// 挨拶用のレスポンス構造体
type GreetingResponse struct {
	Response string `json:"response"`
}

// ABTestHandler は設定に基づいてグリーティングハンドラーを提供します
type ABTestHandler struct {
	config    *config.Config
	methodMap map[string]func(*gin.Context)
}

// NewABTestHandler creates a new ABTestHandler
func NewABTestHandler(cfg *config.Config) *ABTestHandler {
	h := &ABTestHandler{
		config: cfg,
		methodMap: map[string]func(*gin.Context){
			"GetGreeting":           GetGreeting,
			"GetGreetingWithFrench": GetGreetingWithFrench,
		},
	}
	return h
}

// HandleGreeting は設定に基づいてABテストを実行します
func (h *ABTestHandler) HandleGreeting(c *gin.Context) {
	methods := h.config.Greeting.Methods

	totalWeight := 0.0
	for _, method := range methods {
		totalWeight += method.Weight
	}

	r := rand.Float64() * totalWeight

	currentWeight := 0.0
	for _, method := range methods {
		currentWeight += method.Weight
		if r <= currentWeight {
			if handler, exists := h.methodMap[method.Name]; exists {
				handler(c)
				return
			}
		}
	}

	GetGreeting(c)
}

// GetInfo は /info エンドポイントのハンドラー
func GetInfo(c *gin.Context) {
	response := InfoResponse{
		Version: "1.0.0",
	}
	c.JSON(http.StatusOK, response)
}

// GetGreeting は /greeting エンドポイントのハンドラー
func GetGreeting(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"code":   404,
			"error":  "name parameter is required",
		})
		return
	}

	response := GreetingResponse{
		Response: "Hello, " + name,
	}
	c.JSON(http.StatusOK, response)
}

func GetGreetingWithFrench(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"code":   404,
			"error":  "name parameter is required",
		})
		return
	}

	response := GreetingResponse{
		Response: "Bonjour, " + name,
	}
	c.JSON(http.StatusOK, response)
}
