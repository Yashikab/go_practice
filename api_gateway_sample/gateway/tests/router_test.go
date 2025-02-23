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
	// Create mock backend server
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("name") != "Test" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello, Test!"}`))
	}))
	defer backend.Close()

	cfg := &config.Config{
		Endpoints: []config.Endpoint{
			{Path: "/greeting/en", Weight: 1.0},
		},
		BackendURL: backend.URL,
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
