package main

import (
	"github.com/ALZEE23/ApiGo/database"
	"github.com/ALZEE23/ApiGo/handlers"
	"github.com/ALZEE23/ApiGo/middlewares"
	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/", handlers.Test)
		api.GET("/ping", handlers.Ping)
		api.POST("/token", handlers.GenerateToken)
		api.POST("/user/register", handlers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", handlers.Ping)
		}
	}
	return router
}

func main() {
	database.ConnectDb()
	app := setupRoutes()

	app.Run(":3000")
}
