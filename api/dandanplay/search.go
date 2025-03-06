package dandanplay

import (
	"log"
	"run/ikaros/danmuku/config"

	"github.com/go-resty/resty/v2"
)

var Client *resty.Client
var BaseUrl string = "https://api.dandanplay.net"

type SearchEpisodeDetails struct {
	EpisodeId    int    `json:"episodeId"`
	EpisodeTitle string `json:"episodeTitle"`
}
type SearchEpisodesAnime struct {
	AnimeId         int                    `json:"animeId"`
	AnimeTitle      string                 `json:"animeTitle"`
	Type            string                 `json:"type"`
	TypeDescription string                 `json:"typeDescription"`
	Episodes        []SearchEpisodeDetails `json:"episodes"`
}
type SearchEpisodesResponse struct {
	ErrorCode    int                   `json:"errorCode"`
	Success      bool                  `json:"success"`
	ErrorMessage string                `json:"errorMessage"`
	HasMore      bool                  `json:"hasMore"`
	Animes       []SearchEpisodesAnime `json:"animes"`
}

func SearchEpisodesWithKeyword(confg config.Config, anime string, episode string) *SearchEpisodesResponse {
	var url = BaseUrl + "/api/v2/search/episodes"
	var appid = confg.Dandanplay.AppId
	var appSecret = confg.Dandanplay.AppSecret
	if Client == nil {
		Client = resty.New()
	}

	Client.Header.Add("X-AppId", appid)
	Client.Header.Add("X-AppSecret", appSecret)

	resp, err := Client.R().
		SetQueryParam("anime", anime).
		SetResult(&SearchEpisodesResponse{}). // 指定响应的解析类型
		Get(url)                              // 替换为实际的 API URL

	// 处理错误
	if err != nil {
		log.Fatalf("Req fail: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode() != 200 {
		log.Fatalf("Rsp status code fail: %s", resp.Status())
	}

	// 获取解析后的结果
	searchEpisodesResponse := resp.Result().(*SearchEpisodesResponse)

	return searchEpisodesResponse
}
