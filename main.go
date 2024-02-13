package main

import (
	url "example.com/todoApi/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", url.GetTodos)
	router.GET("/todos/completed", url.GetCompletedTodos)
	router.GET("/todos/uncompleted", url.GetUnCompletedTodos)
	router.POST("/todos", url.CreateTodo)
	router.PATCH("/todos/status/:id", url.UpdateCompleteStatus)
	router.DELETE("/todos/:id", url.DeleteTodoById)

	router.Run("localhost:8080")

}
