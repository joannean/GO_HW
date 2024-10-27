package todo

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Mu         sync.Mutex // 鎖用於確保線程安全
	GlobalTodo = []TODO{} // 全局待辦事項列表
	DB         *gorm.DB   // GORM 數據庫連接
)

// TODO 表示一個具有唯一 ID 的任務
type TODO struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Task string `json:"task" binding:"required"`
	Done bool   `json:"done"`
}

// 保存 Todos 到文件
func SaveTodos() error {
	file, err := os.Create(config.Database.JsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(GlobalTodo)
}

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

// 其他函數如 initConfig() 和其他處理 Todo 的函數...
