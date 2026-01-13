// pkg/response/response.go
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess   = 0
	CodeFail      = 1000 // 业务错误从 1000 开始
	CodeServerErr = 500  // 系统错误
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"` // omitempty 避免空 data 输出 null
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  "success",
		Data: data,
	})
}

// Fail 业务失败（用 200 状态码，前端友好）
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// Error 系统异常（用 500）
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: CodeServerErr,
		Msg:  msg,
		Data: nil,
	})
}
