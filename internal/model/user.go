// internal/model/user.go
package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // 内置 ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"size:50;unique;not null"`
	Password   string `gorm:"size:255;not null"` // 后面用 bcrypt 加密
	Email      string `gorm:"size:100;unique"`
	Role       string `gorm:"size:20;default:'user'"` // user / admin
	Status     int    `gorm:"default:1"`              // 1正常 0禁用
}
