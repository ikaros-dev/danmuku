package dandanplay

import (
	"log"
	"run/ikaros/danmuku/config"
	"run/ikaros/danmuku/models"
	"run/ikaros/danmuku/utils"

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

func SearchEpisodesWithKeyword(anime string) *SearchEpisodesResponse {
	var url = BaseUrl + "/api/v2/search/episodes"
	var conf = config.Cfg
	var appid = conf.Dandanplay.AppId
	var appSecret = conf.Dandanplay.AppSecret
	if Client == nil {
		Client = resty.New()
	}

	Client.Header.Add("X-AppId", appid)
	Client.Header.Add("X-AppSecret", appSecret)

	resp, err := Client.R().
		SetQueryParam("anime", anime).
		SetResult(&SearchEpisodesResponse{}). // 指定响应的解析类型
		Get(url)                              // 替换为实际的 API URL

	utils.Debug("Request dandanplay SearchEpisodesWithKeyword api for url: " + url + "?anime=" + anime)

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

// @see https://api.dandanplay.net/swagger/index.html#/%E5%BC%B9%E5%B9%95/Comment_GetComment

type CommentsData struct {
	Cid int    `json:"cid"`
	P   string `json:"p"` // 弹幕参数
	M   string `json:"m"` // 弹幕内容
}

type CommentResponseV2 struct {
	Count    int            `json:"count"`
	Comments []CommentsData `json:"comments"`
}

func GetCommentsWithEpisodeId(episodeId string) *CommentResponseV2 {
	var url = BaseUrl + "/api/v2/comment/" + episodeId
	var conf = config.Cfg
	var appid = conf.Dandanplay.AppId
	var appSecret = conf.Dandanplay.AppSecret
	if Client == nil {
		Client = resty.New()
	}

	Client.Header.Add("X-AppId", appid)
	Client.Header.Add("X-AppSecret", appSecret)

	resp, err := Client.R().
		SetResult(&CommentResponseV2{}). // 指定响应的解析类型
		Get(url)                         // 替换为实际的 API URL

	utils.Debug("Request dandanplay GetCommentsWithEpisodeId api for url: " + url)

	// 处理错误
	if err != nil {
		log.Fatalf("Req fail: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode() != 200 {
		log.Fatalf("Rsp status code fail: %s", resp.Status())
	}

	// 获取解析后的结果
	commentResponseV2 := resp.Result().(*CommentResponseV2)
	return commentResponseV2
}

// @see https://api.dandanplay.net/swagger/index.html#/%E7%95%AA%E5%89%A7/Bangumi_GetBangumiDetailsByBgmtvSubjectId
type BangumiDetailsResponse struct {
	ErrorCode    int                    `json:"errorCode"`
	Success      bool                   `json:"success"`
	ErrorMessage string                 `json:"errorMessage"`
	Bangumi      *models.BangumiDetails `json:"bangumi"`
}

func BangumiGetBangumiDetailsByBgmtvSubjectId(bgmtvSubjectId string) *BangumiDetailsResponse {
	var url = BaseUrl + "/api/v2/bangumi/bgmtv/" + bgmtvSubjectId
	var conf = config.Cfg
	var appid = conf.Dandanplay.AppId
	var appSecret = conf.Dandanplay.AppSecret
	if Client == nil {
		Client = resty.New()
	}

	Client.Header.Add("X-AppId", appid)
	Client.Header.Add("X-AppSecret", appSecret)

	resp, err := Client.R().
		SetResult(&BangumiDetailsResponse{}). // 指定响应的解析类型
		Get(url)                              // 替换为实际的 API URL

	utils.Debug("Request dandanplay BangumiGetBangumiDetailsByBgmtvSubjectId api for url: " + url)

	// 处理错误
	if err != nil {
		log.Fatalf("Req fail: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode() != 200 {
		log.Fatalf("Rsp status code fail: %s", resp.Status())
	}

	// 获取解析后的结果
	bangumiDetailsResponse := resp.Result().(*BangumiDetailsResponse)
	return bangumiDetailsResponse
}
