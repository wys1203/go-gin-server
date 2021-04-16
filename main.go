package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engin := gin.Default()
	engin.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"key": "value"})
	})
	engin.Run(":8888")
}
