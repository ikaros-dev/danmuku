package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchEpisodesWithKeyword(c *gin.Context) {
	// var users []models.User
	// config.DB.Find(&users)
	c.JSON(http.StatusOK, "")
}
