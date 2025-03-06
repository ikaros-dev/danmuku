package models

import "gorm.io/gorm"

type Anime struct {
	gorm.Model
	AnimeId         int    `json:"animeId" gorm:"primaryKey"`
	AnimeTitle      string `json:"animeTitle"`
	Type            string `json:"type"`
	TypeDescription string `json:"typeDescription"`
}
