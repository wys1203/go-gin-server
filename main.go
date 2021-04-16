package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engin := gin.Default()
	engin.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	engin.Run(":8888")
}
