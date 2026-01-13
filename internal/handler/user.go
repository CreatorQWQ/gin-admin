// internal/handler/user.go
package handler

import (
	"github.com/CreatorQWQ/gin-admin/internal/service"
	"github.com/CreatorQWQ/gin-admin/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

var User = new(UserHandler)

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

// Register godoc
// @Summary 用户注册
// @Description 注册新用户，返回成功信息（密码会自动哈希）
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body handler.RegisterReq true "注册信息"
// @Success 200 {object} map[string]interface{} "注册成功"
// @Failure 400 {object} map[string]interface{} "参数错误"
// @Failure 500 {object} map[string]interface{} "服务器错误"
// @Router /api/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 1002, "invalid params")
		return
	}

	if err := service.UserSvc.Register(req.Username, req.Password, req.Email); err != nil {
		response.Fail(c, 1003, err.Error())
		return
	}

	response.Success(c, nil)
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary 用户登录
// @Description 用户名密码登录，返回 JWT token
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body handler.LoginReq true "登录凭证"
// @Success 200 {object} map[string]interface{} "登录成功"
// @Failure 400 {object} map[string]interface{} "参数错误"
// @Failure 401 {object} map[string]interface{} "用户名或密码错误"
// @Router /api/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 1002, "invalid params")
		return
	}

	token, err := service.UserSvc.Login(req.Username, req.Password)
	if err != nil {
		response.Fail(c, 1004, err.Error())
		return
	}

	response.Success(c, gin.H{"token": token})
}
