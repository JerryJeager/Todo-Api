package url

import (
	"fmt"
	"strconv"

	"example.com/todoApi/data"
	"example.com/todoApi/models"
	"example.com/todoApi/utils"
	"github.com/gin-gonic/gin"

	"net/http"
)

// GetTodos             godoc
// @Summary      Get Todos array
// @Description  Responds with the list of all Todos as an Array.
// @Tags         Todos
// @Produce      json
// @Success      200  {array}  models.Todo
// @Router       / [get]
func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Todos)
}

// CreateTodo    godoc
// @Summary      Create a New Todo
// @Description  Create a New Todo Object
// @Tags         Todos
// @Produce      json
// @Param        createTodo body models.Todo true "Create a new Todo Object"
// @Success      201 {object}  models.Todo
// @Router       / [post]
func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid format",
		})
		return
	}
	if isUnique := utils.IsIdUnique(newTodo.ID); !isUnique {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Id already exists",
		})
		return
	}

	data.Todos = append(data.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// GetCompletedTodos godoc
// @Summary          Get completed Todos
// @Description      Gets an array of all completed Todos
// @Tags             Todos
// @Produce          json
// @Success          200 {object} models.Todo
// @Router           /completed [get]
func GetCompletedTodos(c *gin.Context) {
	var completedTodos []models.Todo
	for _, t := range data.Todos {
		if t.COMPLETED {
			completedTodos = append(completedTodos, t)
		}
	}

	c.IndentedJSON(http.StatusOK, completedTodos)
}

// GetUnCompletedTodos godoc
// @Summary            Get uncompleted Todos
// @Description        Gets an array of all uncompleted Todos
// @Tags               Todos
// @Produce            json
// @Success            200 {object} models.Todo
// @Router           /uncompleted [get]
func GetUnCompletedTodos(c *gin.Context) {
	var unCompletedTodos []models.Todo
	for _, t := range data.Todos {
		if !t.COMPLETED {
			unCompletedTodos = append(unCompletedTodos, t)
		}
	}

	c.IndentedJSON(http.StatusOK, unCompletedTodos)
}

// UpdateCompleteStatus godoc
// @Summary             Update Todo Status
// @Description         Change the status of a Todo from uncompleted to completed and vice versa
// @Tags                Todos
// @Produce             json
// @Param               id  path  int  true  "update Todo Status by Todo ID"
// @Param               completed body models.TodoCompletedStatus true "change the status of a Todo"
// @Success             201 {object} models.Todo
// @Router              /status/{id} [patch]
func UpdateCompleteStatus(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Id must be an Integer",
		})
		return
	}
	var completeTodoStatus models.TodoCompletedStatus

	if err := c.BindJSON(&completeTodoStatus); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid format",
		})
		return
	}

	for i, t := range data.Todos {
		if t.ID == idInt {
			data.Todos[i].COMPLETED = completeTodoStatus.COMPLETED
			c.IndentedJSON(http.StatusCreated, data.Todos[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "Todo not found",
	})

}

// DeleteTodoByID     godoc
// @Summary           Delete Todo
// @Description       Delete a Todo by it's Todo ID
// @Tags              Todos
// @Produce           json
// @Param             id  path  int  true  "update Todo Status by Todo ID"
// @Success           200 {array} models.Todo
// @Router            /{id}  [delete]
func DeleteTodoById(c *gin.Context) {
	id := c.Param("id")
	var newTodos []models.Todo
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "ID must be an Integer",
		})
		return
	}

	if doesTodoExist := utils.TodoExits(idInt); !doesTodoExist {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "The Todo for the provided Id not found",
		})
		return
	}

	for _, t := range data.Todos {
		if t.ID != idInt {
			newTodos = append(newTodos, t)
		}
	}

	data.Todos = newTodos[:]

	c.IndentedJSON(http.StatusOK, data.Todos)

}
