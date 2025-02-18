package main

import (
	"danmuku/config"
	"danmuku/models"
	"net/http"

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
			"message": "Hello, Gin!",
		})
	})

	// 创建用户
	r.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		config.DB.Create(&user)
		c.JSON(http.StatusCreated, user)
	})

	// 获取所有用户
	r.GET("/users", func(c *gin.Context) {
		var users []models.User
		config.DB.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	// 获取单个用户
	r.GET("/users/:id", func(c *gin.Context) {
		var user models.User
		if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	// 更新用户
	r.PUT("/users/:id", func(c *gin.Context) {
		var user models.User
		if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		config.DB.Save(&user)
		c.JSON(http.StatusOK, user)
	})

	// 删除用户
	r.DELETE("/users/:id", func(c *gin.Context) {
		var user models.User
		if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		config.DB.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	})

	r.Run() // 默认监听 0.0.0.0:8080
}
