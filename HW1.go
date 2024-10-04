package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// TODO 結構體包含唯一的 ID
type TODO struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

// 全局的 Todo 列表和鎖
var (
	globalTodo = []TODO{}
	idCounter  = 1
	mu         sync.Mutex
)

func main() {
	hw1()
}

func hw1() {
	r := gin.Default()

	// 分組路由
	hw1Group := r.Group("/hw1")
	{
		hw1Group.POST("/todolist", createTodo)
		hw1Group.GET("/todolist", getTodos)
		hw1Group.PUT("/todolist/:id", updateTodo)
		hw1Group.DELETE("/todolist/:id", deleteTodo)
	}

	// 啟動伺服器在 :8008 端口
	r.Run(":8008")
}

// 創建新的 Todo 項目
func createTodo(c *gin.Context) {
	var newTodo TODO

	// 綁定 JSON 請求體到 newTodo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 確保 Task 不為空
	if newTodo.Task == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task cannot be empty"})
		return
	}

	// 分配唯一的 ID
	mu.Lock()
	newTodo.ID = idCounter
	idCounter++
	mu.Unlock()

	// 添加到全局 Todo 列表
	mu.Lock()
	globalTodo = append(globalTodo, newTodo)
	mu.Unlock()

	c.JSON(http.StatusCreated, newTodo)
}

// 獲取所有 Todo 項目，並可根據 Done 狀態過濾
func getTodos(c *gin.Context) {
	doneFilter := c.Query("done")
	var filteredTodos []TODO

	mu.Lock()
	defer mu.Unlock()

	if doneFilter == "" {
		// 返回所有 Todo
		filteredTodos = globalTodo
	} else {
		// 解析 doneFilter 為布林值
		done, err := strconv.ParseBool(doneFilter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for done filter"})
			return
		}

		// 過濾 Todo
		for _, todo := range globalTodo {
			if todo.Done == done {
				filteredTodos = append(filteredTodos, todo)
			}
		}
	}

	c.JSON(http.StatusOK, filteredTodos)
}

// 更新指定 ID 的 Todo 項目
func updateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateData TODO
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, todo := range globalTodo {
		if todo.ID == id {
			// 更新字段
			if updateData.Task != "" {
				globalTodo[i].Task = updateData.Task
			}
			globalTodo[i].Done = updateData.Done
			c.JSON(http.StatusOK, globalTodo[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// 刪除指定 ID 的 Todo 項目
func deleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, todo := range globalTodo {
		if todo.ID == id {
			// 刪除 Todo
			globalTodo = append(globalTodo[:i], globalTodo[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
