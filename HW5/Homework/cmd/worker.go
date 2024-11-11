package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/joannean/GO_HW/todo"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "ToDo Worker",
	Long:  "Will Pop a job from the queue and process it.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Worker called")

		// 初始化 Redis 連接
		redisClient := redis.NewClient(&redis.Options{
			Addr: "localhost:6380",
		})

		// 初始化資料庫連接
		todo.InitializeDatabase("todo.db")

		for {
			// 從 Redis 彈出工作資料
			job, err := redisClient.BRPop(context.Background(), 0, "queue").Result()
			if err != nil {
				fmt.Println("Redis 錯誤:", err)
				continue
			}

			// 解析工作資料
			var todoTask todo.TODO
			err = json.Unmarshal([]byte(job[1]), &todoTask)
			if err != nil {
				fmt.Println("解析 JSON 錯誤:", err)
				continue
			}

			// 將工作資料插入到資料庫
			err = todo.DB.Create(&todoTask).Error
			if err != nil {
				fmt.Println("資料庫插入錯誤:", err)
				continue
			}

			fmt.Println("成功插入到資料庫的任務:", todoTask)
		}
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
