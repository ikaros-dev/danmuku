package routes

import (
	"run/ikaros/danmuku/controllers"

	"github.com/gin-gonic/gin"
)

func SetupDandanplayRoutes(r *gin.Engine) {
	r.GET("/api/dandanplay/search/episodes", controllers.SearchEpisodesWithKeyword)
	r.GET("/api/dandanplay/comment/:episodeId", controllers.GetCommentsWithEpisodeId)
	r.GET("/api/dandanplay/v2/bangumi/bgmtv/:bgmtvSubjectId", controllers.GetBangumiWithBgmtvSubjectId)
}
