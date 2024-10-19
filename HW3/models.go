package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TODO 表示一個具有唯一 ID 的任務
type TODO struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Task string `json:"task" binding:"required"`
	Done bool   `json:"done"`
}

var DB *gorm.DB

// InitializeDatabase 連接到資料庫並執行遷移
func InitializeDatabase(dbPath string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("連接資料庫失敗: %v", err)
	}

	// 自動遷移 TODO 模型
	err = DB.AutoMigrate(&TODO{})
	if err != nil {
		log.Fatalf("資料庫遷移失敗: %v", err)
	}
}
