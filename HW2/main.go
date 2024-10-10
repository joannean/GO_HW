package main

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	mu         sync.Mutex
	globalTodo = []TODO{}
	idCounter  uint // 確保這裡定義為 uint
)

func main() {
	r := gin.Default()

	hw1Group := r.Group("/hw1")
	{
		hw1Group.POST("/todolist", createTodo)
		hw1Group.GET("/todolist", getTodos)
		hw1Group.PUT("/todolist/:id", updateTodo)
		hw1Group.DELETE("/todolist/:id", deleteTodo)
	}

	if err := r.Run(":8008"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
