package main

import (
	"run/ikaros/danmuku/config"
	"run/ikaros/danmuku/models"
	"run/ikaros/danmuku/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 连接数据库
	config.ConnectDatabase()

	// 自动迁移模型到数据库表
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default() // 创建默认路由（含日志和恢复中间件）

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 初始化所有路由
	routes.SetupRouters(r)

	r.Run() // 默认监听 0.0.0.0:8080
}
