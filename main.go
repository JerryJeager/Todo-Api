package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	router := gin.Default()

	router.Run("localhost:8080")

}
