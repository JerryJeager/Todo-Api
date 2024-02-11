package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id" binding:"required"`
	NAME      string `json:"name" binding:"required"`
	COMPLETED bool   `json:"completed" binding:"required"`
}

type todoCompletedStatus struct {
	COMPLETED bool `json:"completed" binding:"required"`
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
		return
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

func updateCompleteStatus(c *gin.Context) {
	id := c.Param("id")
	var completeTodoStatus todoCompletedStatus
	// completeTodoStatus = &todos

	if err := c.BindJSON(&completeTodoStatus); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid format",
		})
		return
	}

	for i, t := range todos {
		if t.ID == id {
			// todos[i].COMPLETED = *completeTodoStatus.COMPLETED
			// todo: fix internal server error from bindjson for bool not working for false values for completed
			todos[i].COMPLETED = completeTodoStatus.COMPLETED
			c.IndentedJSON(http.StatusCreated, todos[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "Todo not found",
	})

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/completed", getCompletedTodos)
	router.GET("/todos/uncompleted", getUnCompletedTodos)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/status/:id", updateCompleteStatus)

	router.Run("localhost:8080")

}
