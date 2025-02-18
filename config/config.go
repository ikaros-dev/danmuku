package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
}

type AppConfig struct {
	Port int `yaml:"port"`
}

// LoadConfig 加载配置文件
func LoadConfig(defaultPath, localPath string) (*Config, error) {
	// 加载默认配置文件
	cfg, err := loadConfigFromFile(defaultPath)
	if err != nil {
		return nil, err
	}

	// 加载本地配置文件（如果存在）
	if _, err := os.Stat(localPath); err == nil {
		localCfg, err := loadConfigFromFile(localPath)
		if err != nil {
			log.Printf("Failed to load local config: %v", err)
		} else {
			// 覆盖默认配置
			mergeConfig(cfg, localCfg)
		}
	}

	return cfg, nil
}

// loadConfigFromFile 从文件加载配置
func loadConfigFromFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// mergeConfig 合并配置（用 localCfg 覆盖 cfg）
func mergeConfig(cfg, localCfg *Config) {
	if localCfg.App.Port != 0 {
		cfg.App.Port = localCfg.App.Port
	}
}
