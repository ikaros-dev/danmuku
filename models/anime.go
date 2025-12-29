package models

type Anime struct {
	AnimeId         int    `json:"animeId" gorm:"primaryKey"`
	AnimeTitle      string `json:"animeTitle"`
	Type            string `json:"type"`
	TypeDescription string `json:"typeDescription"`
}
