package routes

import "github.com/gin-gonic/gin"

func SetupRouters(r *gin.Engine) {
	SetupUserRoutes(r)
	SetupDandanplayRoutes(r)
}
