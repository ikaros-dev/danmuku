package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()  // 创建默认路由（含日志和恢复中间件）

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, Gin!",
        })
    })

    r.Run()  // 默认监听 0.0.0.0:8080
}
