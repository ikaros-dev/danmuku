package main

import (
	"fmt"
	"log"
	"run/ikaros/danmuku/config"
	"run/ikaros/danmuku/models"
	"run/ikaros/danmuku/routes"

	"run/ikaros/danmuku/utils"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

var Version = "1.1.5"

func main() {
	logrus.Info("Starting application ", "v", Version, " ...")
	utils.SetLogLevel(utils.DebugLevel)
	// 加载配置文件
	config.LoadConfig()

	// 连接数据库
	config.ConnectDatabase()

	// 自动迁移模型到数据库表
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Anime{})
	config.DB.AutoMigrate(&models.Episode{})

	r := gin.Default() // 创建默认路由（含日志和恢复中间件）

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 初始化所有路由
	routes.SetupRouters(r)

	port := config.Cfg.App.Port
	logrus.Info("Starting server on port in ", port)
	if err := r.Run(":" + fmt.Sprint(port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	logrus.Info("Application has started.")
}
