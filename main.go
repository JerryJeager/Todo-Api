package main

import (
	"log"
	"os"

	url "example.com/todoApi/api"
	_ "example.com/todoApi/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     Todo App API
// @version         1.0
// @description     A TODO APP API in Go using Gin framework.

// @contact.name   Amadi Jerry
// @contact.url    https://twitter.com/Jerry_Jaeger_
// @contact.email  amadijerry823@gmail.com

// @host
// @BasePath  /todos
func main() {
	router := gin.Default()
	router.GET("/todos", url.GetTodos)
	router.GET("/todos/completed", url.GetCompletedTodos)
	router.GET("/todos/uncompleted", url.GetUnCompletedTodos)
	router.POST("/todos", url.CreateTodo)
	router.PATCH("/todos/status/:id", url.UpdateCompleteStatus)
	router.DELETE("/todos/:id", url.DeleteTodoById)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}
