package config

import (
	"log"
	"os"
	"path/filepath"
	"run/ikaros/danmuku/utils"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App        AppConfig        `yaml:"app"`
	Dandanplay DandanplayConfig `yaml:"dandanplay"`
}

type AppConfig struct {
	Port int `yaml:"port"`
}

type DandanplayConfig struct {
	AppId     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

var Cfg Config

// LoadConfig 加载配置文件
func LoadConfig() {
	configPath := filepath.Join(utils.GetUserHomeAppDir(), "config.yaml")
	DoLoadConfig(configPath)
}

func DoLoadConfig(defaultPath string) (*Config, error) {
	// 加载默认配置文件
	Cfg, err := loadConfigFromFile(defaultPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return nil, err
	}

	// 加载本地配置文件（如果存在）
	// if _, err := os.Stat(localPath); err == nil {
	// 	localCfg, err := loadConfigFromFile(localPath)
	// 	if err != nil {
	// 		log.Printf("Failed to load local config: %v", err)
	// 	} else {
	// 		// 覆盖默认配置
	// 		mergeConfig(Cfg, localCfg)
	// 	}
	// }

	return Cfg, nil
}

// loadConfigFromFile 从文件加载配置
func loadConfigFromFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		return nil, err
	}

	return &Cfg, nil
}

// mergeConfig 合并配置（用 localCfg 覆盖 cfg）
// func mergeConfig(cfg, localCfg *Config) {
// 	if localCfg.App.Port != 0 {
// 		cfg.App.Port = localCfg.App.Port
// 	}

// 	if localCfg.Dandanplay.AppId != "" {
// 		cfg.Dandanplay.AppId = localCfg.Dandanplay.AppId
// 	}

// 	if localCfg.Dandanplay.AppSecret != "" {
// 		cfg.Dandanplay.AppSecret = localCfg.Dandanplay.AppSecret
// 	}
// }
