package main

import (
	"net/http"

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

func main() {
	// Ginルーターの初期化
	r := gin.Default()

	// GETメソッドで/infoエンドポイントを設定
	r.GET("/info", func(c *gin.Context) {
		response := InfoResponse{
			Version: "1.0.0", // ハードコードしたバージョン
		}
		c.JSON(http.StatusOK, response)
	})

	// 挨拶を返すエンドポイント
	r.GET("/greeting", func(c *gin.Context) {
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
	})

	// サーバーを8080ポートで起動
	r.Run(":8080")
}
