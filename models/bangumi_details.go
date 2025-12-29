package models

type BangumiDetails struct {
	AnimeId         int              `json:"animeId" gorm:"primaryKey"`
	BangumiId       string           `json:"bangumiId"`
	BgmtvSubjectId  string           `json:"bgmtvSubjectId"`
	AnimeTitle      string           `json:"animeTitle"`
	Type            string           `json:"type"`
	TypeDescription string           `json:"typeDescription"`
	Episodes        []BangumiEpisode `json:"episodes" gorm:"-"`
}
