package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ResponseError 是統一的錯誤回應函數
func ResponseError(c *gin.Context, statusCode int, errorCode string) {
	// 日誌記錄錯誤
	log.Printf("錯誤 [%s]: %s", errorCode, ErrorMessages[errorCode])
	// 發送錯誤回應
	c.JSON(statusCode, gin.H{"error": errorCode, "message": ErrorMessages[errorCode]})
}

// 保存 Todos 到文件
func saveTodos() error {
	file, err := os.Create("todos.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(globalTodo)
}

// 創建新的 Todo 項目
func createTodo(c *gin.Context) {
	var newTodo TODO

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		ResponseError(c, http.StatusBadRequest, Error6000)
		return
	}

	if newTodo.Task == "" {
		ResponseError(c, http.StatusBadRequest, Error6001)
		return
	}

	mu.Lock() // 獲取鎖
	newTodo.ID = idCounter
	idCounter++
	globalTodo = append(globalTodo, newTodo)
	mu.Unlock() // 釋放鎖

	if err := saveTodos(); err != nil {
		ResponseError(c, http.StatusInternalServerError, Error6002)
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

// 獲取所有 Todo 項目，並可根據 Done 狀態過濾
func getTodos(c *gin.Context) {
	doneFilter := c.Query("done")
	var todos []TODO

	mu.Lock()         // 獲取鎖
	defer mu.Unlock() // 在函數結束時釋放鎖

	if doneFilter == "" {
		todos = globalTodo
	} else {
		done, err := strconv.ParseBool(doneFilter)
		if err != nil {
			ResponseError(c, http.StatusBadRequest, Error6000)
			return
		}

		for _, todo := range globalTodo {
			if todo.Done == done {
				todos = append(todos, todo)
			}
		}
	}

	c.JSON(http.StatusOK, todos)
}

// 更新指定 ID 的 Todo 項目
func updateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, Error6000)
		return
	}

	var updateData TODO
	if err := c.ShouldBindJSON(&updateData); err != nil {
		ResponseError(c, http.StatusBadRequest, Error6000)
		return
	}

	mu.Lock()         // 獲取鎖
	defer mu.Unlock() // 在函數結束時釋放鎖

	for i, todo := range globalTodo {
		if todo.ID == uint(id) { // 將 id 轉換為 uint
			if updateData.Task != "" {
				globalTodo[i].Task = updateData.Task
			}
			globalTodo[i].Done = updateData.Done
			if err := saveTodos(); err != nil {
				ResponseError(c, http.StatusInternalServerError, Error6004)
				return
			}
			c.JSON(http.StatusOK, globalTodo[i])
			return
		}
	}

	ResponseError(c, http.StatusNotFound, Error6005)
}

// 刪除指定 ID 的 Todo 項目
func deleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, Error6000)
		return
	}

	mu.Lock()         // 獲取鎖
	defer mu.Unlock() // 在函數結束時釋放鎖

	for i, todo := range globalTodo {
		if todo.ID == uint(id) { // 將 id 轉換為 uint
			globalTodo = append(globalTodo[:i], globalTodo[i+1:]...)
			if err := saveTodos(); err != nil {
				ResponseError(c, http.StatusInternalServerError, Error6006)
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	ResponseError(c, http.StatusNotFound, Error6005)
}
