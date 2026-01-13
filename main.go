package main

import (
	"github.com/CreatorQWQ/gin-admin/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		pkg.Fail(c, 500, "internal error")
	}))
	r.GET("/ping", func(c *gin.Context) {
		pkg.Success(c, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
