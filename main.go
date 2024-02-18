package main

import (
	// "database/sql"
	// "fmt"
	"log"
	"os"

	"example.com/todoApi/config"
	_ "github.com/lib/pq"

	url "example.com/todoApi/api"
	_ "example.com/todoApi/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// var db *sql.DB

// const (
// 	host     = "localhost"
// 	port     = "5433"
// 	user     = "postgres"
// 	password = "chidiebere823"
// 	dbname   = "TODO_DB"
// )

// func init() {

// 	var err error

// 	connStr := "postgres://postgres:password@localhost/DB_1?sslmode=disable"
// 	db, err = sql.Open("postgres", connStr)

// 	if err != nil{
// 		panic(err)
// 	}

// 	if err := db.Ping(); err != nil{
// 		panic(err)
// 	}

// 	fmt.Println("The database is connected.")

// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// @title     Todo App API
// @version         1.0
// @description     A TODO APP API in Go using Gin framework.

// @contact.name   Amadi Jerry
// @contact.url    https://twitter.com/Jerry_Jaeger_
// @contact.email  amadijerry823@gmail.com

// @host
// @BasePath  /todos
func main() {

	config.Connect()

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
