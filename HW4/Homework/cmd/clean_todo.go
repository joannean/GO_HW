package cmd

import (
	"fmt"
	"log"

	"github.com/joannean/GO_HW/todo"
	"github.com/spf13/cobra"
)

// cleanTodoCmd represents the cleanTodo command
var cleanTodoCmd = &cobra.Command{
	Use:   "clean_todo",
	Short: "Delete completed todos",
	Long:  `This command deletes all todos that are marked as done (done == true).`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := deleteDoneTodos(); err != nil {
			log.Fatalf("Error cleaning todos: %v", err)
		}
		fmt.Println("Successfully deleted completed todos.")
	},
}

// deleteDoneTodos deletes all todos that are marked as done
func deleteDoneTodos() error {
	// 使用全局鎖確保數據一致性
	todo.Mu.Lock()
	defer todo.Mu.Unlock()

	var remainingTodos []todo.TODO
	for _, t := range todo.GlobalTodo {
		if !t.Done { // 如果未完成則保留
			remainingTodos = append(remainingTodos, t)
		}
	}

	// 更新全局 todo 列表
	todo.GlobalTodo = remainingTodos

	// 保存更新的 todo 列表
	if err := todo.SaveTodos(); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(cleanTodoCmd)
}
