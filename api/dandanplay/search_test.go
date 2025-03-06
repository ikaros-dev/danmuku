package dandanplay

import (
	"run/ikaros/danmuku/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchEpisodesWithKeyword(t *testing.T) {
	t.Skip("This test is disabled")
	config.DoLoadConfig("../../config.yaml", "../../config.local.yaml")

	// 测试搜索功能
	keyword := "CLANNAD" // 替换为你想要搜索的关键字
	resp := SearchEpisodesWithKeyword(keyword)

	assert.EqualValues(t, 0, resp.ErrorCode)

	// assert.NotEmpty(t, resp)

	// // 打印返回的数据以便检查
	// for _, episode := range episodes {
	// 	log.Printf("Found episode: %+v\n", episode)
	// }
}
