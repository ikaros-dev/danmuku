package main

import (
	"fmt"
	"log"
	"run/ikaros/danmuku/config"
	"run/ikaros/danmuku/models"
	"run/ikaros/danmuku/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

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

	port := cfg.App.Port
	log.Printf("Starting server on port %d...\n", port)
	if err := r.Run(":" + fmt.Sprint(port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
