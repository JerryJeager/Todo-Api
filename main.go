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

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)

	router.Run("localhost:8080")

}
