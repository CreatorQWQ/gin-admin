// main.go
package main

import (
	"github.com/CreatorQWQ/gin-admin/internal/handler"
	"github.com/CreatorQWQ/gin-admin/internal/middleware"
	"github.com/CreatorQWQ/gin-admin/internal/model"
	"github.com/CreatorQWQ/gin-admin/pkg/common"
	"github.com/CreatorQWQ/gin-admin/pkg/response"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB() // 初始化数据库
	common.InitRedis()

	// 自动迁移（开发时用，生产慎用或用迁移工具）
	common.DB.AutoMigrate(&model.User{})
	common.DB.AutoMigrate(&model.Article{})

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
		api.POST("/register", handler.User.Register)
		api.POST("/login", handler.User.Login)

		// 示例保护路由（测试 auth 中间件）
		api.GET("/profile", middleware.Auth(), func(c *gin.Context) {
			userID := c.GetUint("user_id")
			response.Success(c, gin.H{"user_id": userID, "msg": "protected route"})

		})
		api.POST("/articles", middleware.Auth(), handler.Article.Create)
		api.GET("/articles", middleware.Auth(), handler.Article.List) // 可根据需求去掉 Auth 开放列表
		api.PUT("/articles/:id", middleware.Auth(), handler.Article.Update)
		api.DELETE("/articles/:id", middleware.Auth(), handler.Article.Delete)
	}

	r.Run(":8080") // 或从配置读端口
}
