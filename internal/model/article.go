// internal/model/article.go
package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string `gorm:"size:200;not null"`
	Content  string `gorm:"type:text;not null"`
	AuthorID uint   `gorm:"not null"`  // 关联用户ID
	Status   int    `gorm:"default:1"` // 1:正常 0:草稿/删除
}
