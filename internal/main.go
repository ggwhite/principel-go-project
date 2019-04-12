package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.LoadHTMLGlob("web/*")

	e.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404", gin.H{})
	})

	e.GET("", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hello World",
		})
	})

	e.Run(":3000")
}
