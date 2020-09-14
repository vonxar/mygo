package main

import (
	// WAFのGinをインポート
	"github.com/gin-gonic/gin"
)

func main() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	router := gin.Default()

	// ルーター設定
	// ブラウザで「/」 にアクセスしたら「Hello Gin!」と表示される設定です
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!")
	})
	router.Run(":8080")
}
