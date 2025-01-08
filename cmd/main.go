package main

import (
	"time"

	"github.com/ALZEE23/ApiGo/database"
	"github.com/ALZEE23/ApiGo/handlers"
	"github.com/ALZEE23/ApiGo/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173", "https://web.pplg-game.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))
	router.Static("/storage", "./storage")
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
