package dandanplay

import (
	"os"
	"run/ikaros/danmuku/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchEpisodesWithKeyword(t *testing.T) {
	t.Skip("This test is disabled")
	// 从环境变量中获取 AppId 和 AppSecret
	appId := os.Getenv("DANDANPLAY_APP_ID")
	appSecret := os.Getenv("DANDANPLAY_APP_SECRET")

	// 如果环境变量未设置，跳过测试
	if appId == "" || appSecret == "" {
		t.Skip("DANDANPLAY_APP_ID or DANDANPLAY_APP_SECRET environment variables not set")
	}

	// 初始化配置
	conf := config.Config{
		Dandanplay: config.DandanplayConfig{
			AppId:     appId,
			AppSecret: appSecret,
		},
	}

	// 测试搜索功能
	keyword := "CLANNAD" // 替换为你想要搜索的关键字
	resp := SearchEpisodesWithKeyword(conf, keyword, "")

	// 至少返回一条数据
	assert.EqualValues(t, 0, resp.ErrorCode)

	// assert.NotEmpty(t, resp)

	// // 打印返回的数据以便检查
	// for _, episode := range episodes {
	// 	log.Printf("Found episode: %+v\n", episode)
	// }
}
