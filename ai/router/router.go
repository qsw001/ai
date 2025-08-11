package router

import (
	"ai/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	// 跨域配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 注册路由
	r.POST("/chat", func(c *gin.Context) {
		handler.ChatHandler(c.Writer, c.Request)
	})

	return r
}
