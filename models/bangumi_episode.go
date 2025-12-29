package models

type BangumiEpisode struct {
	EpisodeId      int    `json:"episodeId" gorm:"primaryKey"`
	AnimeId        int    `json:"animeId"`
	BgmtvSubjectId string `json:"bgmtvSubjectId"`
	SeasonId       int    `json:"seasonId"`
	EpisodeTitle   string `json:"episodeTitle"`
	EpisodeNumber  string `json:"episodeNumber"`
	LastWatched    string `json:"lastWatched"`
	AirDate        string `json:"airDate"`
}
