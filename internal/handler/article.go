package handler

import (
	"strconv"

	"github.com/CreatorQWQ/gin-admin/internal/service"
	"github.com/CreatorQWQ/gin-admin/pkg/response"
	"github.com/gin-gonic/gin"
)

var Article = new(ArticleHandler)

type ArticleHandler struct{}

type CreateReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var req CreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 1002, "参数错误")
		return
	}

	userID := c.GetUint("user_id")
	if err := service.ArticleSvc.Create(req.Title, req.Content, userID); err != nil {
		response.Fail(c, 1005, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ArticleHandler) List(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")
	keyword := c.Query("keyword")

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	articles, total, err := service.ArticleSvc.List(page, size, keyword)
	if err != nil {
		response.Fail(c, 1005, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  articles,
		"total": total,
	})
}

func (h *ArticleHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Fail(c, 1002, "ID 无效")
		return
	}

	var req UpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 1002, "参数错误")
		return
	}

	userID := c.GetUint("user_id")
	if err := service.ArticleSvc.Update(uint(id), req.Title, req.Content, userID); err != nil {
		response.Fail(c, 1005, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Fail(c, 1002, "ID 无效")
		return
	}

	userID := c.GetUint("user_id")
	if err := service.ArticleSvc.Delete(uint(id), userID); err != nil {
		response.Fail(c, 1005, err.Error())
		return
	}
	response.Success(c, nil)
}
