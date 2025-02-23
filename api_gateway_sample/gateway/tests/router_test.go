package tests

import (
	"gateway/config"
	"gateway/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRouter(t *testing.T) {
	cfg := &config.Config{
		Endpoints: []config.Endpoint{
			{Path: "/greeting/en", Weight: 1.0},
		},
		BackendURL: "http://localhost:8080",
	}

	router := handlers.NewRouter(cfg)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("POST", "/greeting", nil)
	req.Header.Set("name", "Test")
	c.Request = req

	router.RouteGreeting(c)

	if w.Code != http.StatusOK {
		t.Errorf("want status code %d, got %d", http.StatusOK, w.Code)
	}
}
