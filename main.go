package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        int    `json:"id" binding:"required"`
	NAME      string `json:"name" binding:"required"`
	COMPLETED bool   `json:"completed" binding:"required"`
}

type todoCompletedStatus struct {
	COMPLETED bool `json:"completed" binding:"required"`
}

var todos = []todo{
	{ID: 1, NAME: "Wash dirty clothes", COMPLETED: false},
	{ID: 2, NAME: "Watch season 5 of My Hero Acamedia", COMPLETED: false},
	{ID: 3, NAME: "Learn Backend Development with Go", COMPLETED: false},
	{ID: 4, NAME: "Buy Apple Vision Pro", COMPLETED: false},
	{ID: 5, NAME: "Watch season 2 of Jujutsu Kaisen", COMPLETED: true},
}

func isIdUnique(id int) bool {
	for _, t := range todos {
		if t.ID == id {
			return false
		}
	}
	return true
}

func todoExits(id int) bool {
	for _, t := range todos {
		if t.ID == id {
			return true
		}
	}
	return false
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
	if isUnique := isIdUnique(newTodo.ID); !isUnique {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Id already exists",
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
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Id must be an Integer",
		})
		return
	}
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
		if t.ID == idInt {
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

func deleteTodoById(c *gin.Context) {
	id := c.Param("id")
	var newTodos []todo
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "ID must be an Integer",
		})
		return
	}

	if doesTodoExist := todoExits(idInt); !doesTodoExist {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "The Todo for the provided Id not found",
		})
		return
	}

	for _, t := range todos {
		if t.ID != idInt {
			newTodos = append(newTodos, t)
		}
	}

	todos = newTodos[:]

	c.IndentedJSON(http.StatusOK, todos)

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/completed", getCompletedTodos)
	router.GET("/todos/uncompleted", getUnCompletedTodos)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/status/:id", updateCompleteStatus)
	router.DELETE("/todos/:id", deleteTodoById)

	router.Run("localhost:8080")

}
