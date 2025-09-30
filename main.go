package main

import (
	"github.com/tsutsutsu123/only-login-and-signup/controllers"
	"github.com/tsutsutsu123/only-login-and-signup/models"

	"github.com/gin-gonic/gin"
)
func setupRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)

	return router
}

func main() {
	models.ConnectDataBase()
	router := setupRouter()
	router.Run(":8080")
}
