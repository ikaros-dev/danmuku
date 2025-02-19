package routes

import (
	"run/ikaros/danmuku/controllers"

	"github.com/gin-gonic/gin"
)

func SetupDandanplayRoutes(r *gin.Engine) {
	r.GET("/api/dandanplay/search/episodes", controllers.SearchEpisodesWithKeyword)
}
