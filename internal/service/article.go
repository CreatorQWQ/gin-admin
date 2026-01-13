package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/CreatorQWQ/gin-admin/internal/model"
	"github.com/CreatorQWQ/gin-admin/pkg/common"
)

type ArticleService struct{}

var ArticleSvc = new(ArticleService)

const (
	articleListCachePrefix = "article:list:"
	articleCacheTTL        = 5 * time.Minute
)

func (s *ArticleService) Create(title, content string, authorID uint) error {
	article := model.Article{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	}
	return common.DB.Create(&article).Error
}

func (s *ArticleService) List(page, size int, keyword string) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	ctx := context.Background()

	// 简单缓存 key（实际生产可 hash 参数更精确）
	cacheKey := fmt.Sprintf("%s%d-%d-%s", articleListCachePrefix, page, size, keyword)
	totalKey := cacheKey + ":total"

	// 尝试从 Redis 获取
	val, err := common.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &articles)
		common.Redis.Get(ctx, totalKey).Scan(&total)
		return articles, total, nil
	}

	// DB 查询
	query := common.DB.Model(&model.Article{}).Where("status = 1")
	if keyword != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	query.Count(&total)

	query = query.Order("created_at desc").
		Offset((page - 1) * size).
		Limit(size).
		Find(&articles)

	if err := query.Error; err != nil {
		return nil, 0, err
	}

	// 写入缓存
	data, _ := json.Marshal(articles)
	common.Redis.Set(ctx, cacheKey, data, articleCacheTTL)
	common.Redis.Set(ctx, totalKey, total, articleCacheTTL)

	return articles, total, nil
}

// 简单 Update（示例，可扩展）
func (s *ArticleService) Update(id uint, title, content string, userID uint) error {
	var article model.Article
	if err := common.DB.First(&article, id).Error; err != nil {
		return err
	}
	// 简单权限检查：只能改自己的文章
	if article.AuthorID != userID {
		return fmt.Errorf("permission denied")
	}

	article.Title = title
	article.Content = content
	return common.DB.Save(&article).Error
}

// Delete（软删除）
func (s *ArticleService) Delete(id uint, userID uint) error {
	var article model.Article
	if err := common.DB.First(&article, id).Error; err != nil {
		return err
	}
	if article.AuthorID != userID {
		return fmt.Errorf("permission denied")
	}
	article.Status = 0
	return common.DB.Save(&article).Error
}
