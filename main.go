// main.go
package main

import (
	"github.com/CreatorQWQ/gin-admin/internal/handler"
	"github.com/CreatorQWQ/gin-admin/pkg/response"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 全局 panic 恢复（非常重要，生产级必须有）
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		response.Error(c, "internal server error")
		// 可以在这里加日志：zap.L().Error("panic", zap.Any("recover", recovered))
	}))

	// 路由组（后面会扩展）
	api := r.Group("/api")
	{
		api.GET("/ping", handler.Ping)
	}

	r.Run(":8080") // 或从配置读端口
}
