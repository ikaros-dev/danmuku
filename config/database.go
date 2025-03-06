package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	// 连接 SQLite 数据库（文件路径为 "./data/data.db"）
	DB, err = gorm.Open(sqlite.Open("data/data.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	println("Database connected successfully!")
}
