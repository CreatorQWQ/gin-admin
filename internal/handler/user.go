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
