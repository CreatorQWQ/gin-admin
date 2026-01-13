// internal/handler/ping.go
package handler

import (
	"github.com/CreatorQWQ/gin-admin/pkg/response" // 注意：用你的实际模块名
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response.Success(c, map[string]string{"message": "pong"})
}
