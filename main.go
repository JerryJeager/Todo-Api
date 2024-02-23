package main

import (
	// "database/sql"
	// "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	// _ "github.com/jinzhu/gorm/dialects/postgres"

	url "example.com/todoApi/api"
	"example.com/todoApi/config"
	_ "example.com/todoApi/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Connect()
}

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
	router.POST("/user", url.CreateUser)
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
