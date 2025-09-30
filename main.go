package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "this is the register endpoint.",
		})
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
