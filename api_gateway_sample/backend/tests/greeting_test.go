package tests

import (
	"backend/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGreetings(t *testing.T) {
	tests := []struct {
		name       string
		handler    gin.HandlerFunc
		wantBody   string
		headerName string
	}{
		{
			name:       "English greeting",
			handler:    handlers.GreetingEnglish,
			wantBody:   `{"message":"Hello, Test!"}`,
			headerName: "Test",
		},
		{
			name:       "French greeting",
			handler:    handlers.GreetingFrench,
			wantBody:   `{"message":"Bonjour, Test!"}`,
			headerName: "Test",
		},
		{
			name:       "Spanish greeting",
			handler:    handlers.GreetingSpanish,
			wantBody:   `{"message":"¡Hola, Test!"}`,
			headerName: "Test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req := httptest.NewRequest("POST", "/", nil)
			req.Header.Set("name", tt.headerName)
			c.Request = req

			tt.handler(c)

			if w.Code != http.StatusOK {
				t.Errorf("want status code %d, got %d", http.StatusOK, w.Code)
			}
			if w.Body.String() != tt.wantBody {
				t.Errorf("want body %s, got %s", tt.wantBody, w.Body.String())
			}
		})
	}
}
