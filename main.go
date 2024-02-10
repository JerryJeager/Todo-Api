package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	NAME      string `json:"name"`
	COMPLETED bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", NAME: "Wash dirty clothes", COMPLETED: false},
	{ID: "2", NAME: "Watch season 5 of My Hero Acamedia", COMPLETED: false},
	{ID: "3", NAME: "Learn Backend Development with Go", COMPLETED: false},
	{ID: "4", NAME: "Buy Apple Vision Pro", COMPLETED: false},
	{ID: "5", NAME: "Watch season 2 of Jujutsu Kaisen", COMPLETED: true},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid format",
		})
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getCompletedTodos(c *gin.Context) {
	var completedTodos []todo
	for _, t := range todos {
		if t.COMPLETED {
			completedTodos = append(completedTodos, t)
		}
	}

	c.IndentedJSON(http.StatusOK, completedTodos)
}

func getUnCompletedTodos(c *gin.Context) {
	var unCompletedTodos []todo
	for _, t := range todos {
		if !t.COMPLETED {
			unCompletedTodos = append(unCompletedTodos, t)
		}
	}

	c.IndentedJSON(http.StatusOK, unCompletedTodos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/completed", getCompletedTodos)
	router.GET("/todos/uncompleted", getUnCompletedTodos)
	router.POST("/todos", createTodo)

	router.Run("localhost:8080")

}
