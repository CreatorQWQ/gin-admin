// pkg/common/db.go
package common

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:mysql5498138.@tcp(127.0.0.1:3306)/gin_admin?charset=utf8mb4&parseTime=True&loc=Local" // 修改成你的用户名:密码@host:port/数据库名

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 开发时看 SQL 日志，生产可改 Silent
	})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connected successfully!")
}
