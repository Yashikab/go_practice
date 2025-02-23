package main

import (
	"api_practice/config"
	"api_practice/handlers"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 乱数生成器の初期化（新しい方法）
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 設定ファイルを読み込む
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Ginルーターの初期化
	r := gin.Default()

	// ABテストハンドラーの初期化
	abHandler := handlers.NewABTestHandler(cfg)

	// ルーティングの設定
	r.GET("/info", handlers.GetInfo)
	r.GET("/greeting", abHandler.HandleGreeting)

	// サーバーを8080ポートで起動
	r.Run(":8080")
}
