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
		api.POST("/token", handlers.GenerateToken)
		api.POST("/user/register", handlers.RegisterUser)
		api.POST("/apk", handlers.Apk)
		api.GET("/apk", handlers.GetApk)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", handlers.Ping)
			secured.POST("/apk", handlers.Apk)
		}
	}
	return router
}

func main() {
	database.ConnectDb()
	app := setupRoutes()

	app.Run(":3000")
}
