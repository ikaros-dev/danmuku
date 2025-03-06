package models

import "gorm.io/gorm"

type KeywordAnime struct {
	gorm.Model
	AnimeId int    `json:"animeId"`
	Keyword string `json:"keyword"`
}
