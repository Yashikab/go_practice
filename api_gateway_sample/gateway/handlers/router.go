package handlers

import (
	"context"
	"gateway/config"
	"io"
	"math/rand"
	"net/http"
	"time"

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
	// コンテキストの作成（タイムアウト付き）
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel() // タイムアウト用のコンテキストを確実にキャンセル

	// 非同期でバックエンドにリクエストを送信
	responseChan := make(chan *http.Response)
	errorChan := make(chan error)

	go func() {
		// バックエンドへのリクエスト処理
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

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.config.BackendURL+selectedEndpoint, nil)
		if err != nil {
			errorChan <- err
			return
		}

		req.Header = c.Request.Header
		resp, err := r.client.Do(req)
		if err != nil {
			errorChan <- err
			return
		}

		responseChan <- resp
	}()

	// レスポンスを待機
	select {
	case <-ctx.Done():
		// タイムアウトまたはキャンセル
		c.JSON(http.StatusGatewayTimeout, gin.H{"error": "request timeout"})
		return
	case err := <-errorChan:
		// エラー発生
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	case resp := <-responseChan:
		defer resp.Body.Close() // レスポンスボディを確実にクローズ
		// 正常なレスポンス
		c.Status(resp.StatusCode)
		for k, v := range resp.Header {
			c.Writer.Header()[k] = v
		}
		if _, err := io.Copy(c.Writer, resp.Body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to copy response body"})
			return
		}
	}
}
