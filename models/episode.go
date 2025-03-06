package models

import "gorm.io/gorm"

type Episode struct {
	gorm.Model
	AnimeId      int    `json:"animeId"`
	EpisodeId    int    `json:"episodeId" gorm:"primaryKey"`
	EpisodeTitle string `json:"episodeTitle"`
}
