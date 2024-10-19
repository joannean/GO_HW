package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	mu         sync.Mutex
	globalTodo = []TODO{}
	idCounter  uint
	config     Config
)

// Config 定義從配置文件讀取的結構
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Path     string `mapstructure:"path"`
		JsonFile string `mapstructure:"jsonFile"`
	} `mapstructure:"database"`
}

// 初始化配置
func initConfig() {
	// 設定viper來讀取config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // 當前目錄

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("讀取配置文件失敗: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("解析配置文件失敗: %v", err)
	}
}

func main() {
	initConfig() // 初始化配置

	// 初始化資料庫
	InitializeDatabase(config.Database.Path)

	r := gin.Default()

	hw1Group := r.Group("/todoapi")
	{
		hw1Group.POST("/todolist", createTodo)
		hw1Group.GET("/todolist", getTodos)
		hw1Group.PUT("/todolist/:id", updateTodo)
		hw1Group.DELETE("/todolist/:id", deleteTodo)
	}

	port := fmt.Sprintf(":%d", config.Server.Port)
	if err := r.Run(port); err != nil {
		log.Fatalf("無法啟動伺服器: %v", err)
	}
}
