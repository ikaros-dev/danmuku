package config

import (
	"path/filepath"
	"run/ikaros/danmuku/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dbPath := filepath.Join(utils.GetUserHomeAppDir(), "data.db")
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	println("Database connected successfully!")
}
