package controllers

import (
	"log"
	"net/http"
	"run/ikaros/danmuku/api/dandanplay"
	"run/ikaros/danmuku/config"
	"run/ikaros/danmuku/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnimeRsp struct {
	gorm.Model
	AnimeId         int              `json:"animeId"`
	AnimeTitle      string           `json:"animeTitle"`
	Type            string           `json:"type"`
	TypeDescription string           `json:"typeDescription"`
	Episodes        []models.Episode `json:"episodes"`
}

func searchFirstAnim(title string) []AnimeRsp {
	var myAnimes []models.Anime
	result := config.DB.Where("anime_title LIKE ?", "%"+title+"%").Find(&myAnimes)
	if result.Error == nil {
		// 如果查询到了结果，直接组装带有episode的anime返回
		var myAnimeRsps []AnimeRsp
		for _, myAnime := range myAnimes {
			var episodes []models.Episode
			config.DB.Where("anime_id = ?", myAnime.AnimeId).Find(&episodes)
			myAnimeRsps = append(myAnimeRsps, AnimeRsp{
				AnimeId:         myAnime.AnimeId,
				AnimeTitle:      myAnime.AnimeTitle,
				Type:            myAnime.Type,
				TypeDescription: myAnime.AnimeTitle,
				Episodes:        episodes,
			})
		}
		return myAnimeRsps
	}
	return nil
}

func SearchEpisodesWithKeyword(c *gin.Context) {
	anime := c.Query("anime")
	if anime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'anime' parameter"})
		return
	}

	// 先搜索数据库，如存在直接响应
	animeRsps := searchFirstAnim(anime)
	if animeRsps != nil {
		c.JSON(http.StatusOK, gin.H{
			"animes": animeRsps,
		})
		return
	}

	// 如不存在则去查dandanplay，存到数据库并响应
	searchAnimeRsp := dandanplay.SearchEpisodesWithKeyword(anime)

	if len(searchAnimeRsp.Animes) == 0 {
		log.Println("No anime found in search result")
		c.JSON(http.StatusNotFound, "")
		return
	}

	// 开启事务
	tx := config.DB.Begin()
	if tx.Error != nil {
		log.Println("Failed to begin transaction:", tx.Error)
		return
	}
	// 确保在发生错误时回滚事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Transaction rolled back due to panic:", r)
		}
	}()
	// 执行数据库操作
	for _, seAnime := range searchAnimeRsp.Animes {
		anime := &models.Anime{
			AnimeId:         seAnime.AnimeId,
			AnimeTitle:      seAnime.AnimeTitle,
			Type:            seAnime.Type,
			TypeDescription: seAnime.TypeDescription,
		}
		if err := tx.Save(anime).Error; err != nil {
			tx.Rollback()
			log.Println("Failed to save anime:", err)
			return
		}
		for _, seEpisode := range seAnime.Episodes {
			episode := &models.Episode{
				AnimeId:      anime.AnimeId,
				EpisodeId:    seEpisode.EpisodeId,
				EpisodeTitle: seEpisode.EpisodeTitle,
			}
			if err := tx.Save(episode).Error; err != nil {
				tx.Rollback()
				log.Println("Failed to save episode:", err)
				return
			}
		}
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Println("Failed to commit transaction:", err)
		return
	}

	newAnimeRsps := searchFirstAnim(anime)
	c.JSON(http.StatusOK, gin.H{
		"animes": newAnimeRsps,
	})
}

func GetCommentsWithEpisodeId(c *gin.Context) {
	episodeId := c.Param("episodeId")
	commentRsps := dandanplay.GetCommentsWithEpisodeId(episodeId)
	c.JSON(http.StatusOK, gin.H{
		"comments": commentRsps.Comments,
		"count":    commentRsps.Count,
	})
	return
}
