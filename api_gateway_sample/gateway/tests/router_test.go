package tests

import (
	"gateway/config"
	"gateway/handlers"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestRouter_Success(t *testing.T) {
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

	expectedBody := `{"message":"Hello, Test!"}`
	if strings.TrimSpace(w.Body.String()) != expectedBody {
		t.Errorf("want body %s, got %s", expectedBody, w.Body.String())
	}
}

func TestRouter_Timeout(t *testing.T) {
	// Create slow mock backend server
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Delay response longer than context timeout
		time.Sleep(6 * time.Second)
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

	if w.Code != http.StatusGatewayTimeout {
		t.Errorf("want status code %d, got %d", http.StatusGatewayTimeout, w.Code)
	}
}

func TestRouter_BackendError(t *testing.T) {
	// Create backend server that always returns error
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Internal Server Error")
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

	if w.Code != http.StatusInternalServerError {
		t.Errorf("want status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}
