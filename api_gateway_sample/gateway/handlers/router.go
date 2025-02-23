package handlers

import (
	"gateway/config"
	"io"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	config *config.Config
	client *http.Client
}

func NewRouter(config *config.Config) *Router {
	return &Router{
		config: config,
		client: &http.Client{},
	}
}

func (r *Router) RouteGreeting(c *gin.Context) {
	randomValue := rand.Float64()
	var cumulativeWeight float64
	var selectedEndpoint string

	for _, endpoint := range r.config.Endpoints {
		cumulativeWeight += endpoint.Weight
		if randomValue <= cumulativeWeight {
			selectedEndpoint = endpoint.Path
			break
		}
	}

	req, err := http.NewRequest(http.MethodPost, r.config.BackendURL+selectedEndpoint, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Copy headers from original request
	req.Header = c.Request.Header

	resp, err := r.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Copy response status and headers
	c.Status(resp.StatusCode)
	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}

	// Copy response body
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to copy response body"})
		return
	}
}
