package tests

import (
	"api_practice/config"
	"api_practice/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	// テスト用のGinコンテキストを設定
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// ハンドラーを実行
	handlers.GetInfo(c)

	// レスポンスを検証
	assert.Equal(t, http.StatusOK, w.Code)

	var response handlers.InfoResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "1.0.0", response.Version)
}

func TestGetGreeting(t *testing.T) {
	tests := []struct {
		name          string
		queryName     string
		expectedCode  int
		expectedBody  string
		expectedError bool
	}{
		{
			name:         "正常なリクエスト",
			queryName:    "テスト",
			expectedCode: http.StatusOK,
			expectedBody: "Hello, テスト",
		},
		{
			name:          "名前パラメータなし",
			queryName:     "",
			expectedCode:  http.StatusNotFound,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// クエリパラメータを設定
			if tt.queryName != "" {
				c.Request = httptest.NewRequest("GET", "/?name="+tt.queryName, nil)
			} else {
				c.Request = httptest.NewRequest("GET", "/", nil)
			}

			// ハンドラーを実行
			handlers.GetGreeting(c)

			// ステータスコードを検証
			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedError {
				var errorResponse gin.H
				err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)
				assert.Equal(t, "error", errorResponse["status"])
			} else {
				var response handlers.GreetingResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response.Response)
			}
		})
	}
}

func TestGetGreetingWithFrench(t *testing.T) {
	tests := []struct {
		name          string
		queryName     string
		expectedCode  int
		expectedBody  string
		expectedError bool
	}{
		{
			name:         "正常なリクエスト",
			queryName:    "テスト",
			expectedCode: http.StatusOK,
			expectedBody: "Bonjour, テスト",
		},
		{
			name:          "名前パラメータなし",
			queryName:     "",
			expectedCode:  http.StatusNotFound,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if tt.queryName != "" {
				c.Request = httptest.NewRequest("GET", "/?name="+tt.queryName, nil)
			} else {
				c.Request = httptest.NewRequest("GET", "/", nil)
			}

			handlers.GetGreetingWithFrench(c)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedError {
				var errorResponse gin.H
				err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)
				assert.Equal(t, "error", errorResponse["status"])
			} else {
				var response handlers.GreetingResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response.Response)
			}
		})
	}
}

func TestABTestHandler(t *testing.T) {
	// テスト用の設定を作成
	cfg := &config.Config{
		Greeting: struct {
			Methods []config.GreetingMethod `yaml:"methods"`
		}{
			Methods: []config.GreetingMethod{
				{
					Name:   "GetGreeting",
					Weight: 1.0,
				},
			},
		},
	}

	handler := handlers.NewABTestHandler(cfg)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/?name=テスト", nil)

	handler.HandleGreeting(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response handlers.GreetingResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response.Response, "テスト")
}
