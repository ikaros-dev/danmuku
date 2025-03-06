package utils

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func GetUserHomeAppDir() string {
	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.Error("无法获取用户目录:", err)
		return "" // 或者返回一个默认路径，如当前目录
	}
	// 拼接路径
	appDir := filepath.Join(homeDir, ".ikaros_danmuku")
	// 确保目录存在
	if err := os.MkdirAll(appDir, 0755); err != nil {
		logrus.Error("无法创建应用程序目录:", err)
		return "" // 或者返回一个默认路径
	}
	return appDir
}
