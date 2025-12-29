package models

type Episode struct {
	AnimeId      int    `json:"animeId"`
	EpisodeId    int    `json:"episodeId" gorm:"primaryKey"`
	EpisodeTitle string `json:"episodeTitle"`
}
