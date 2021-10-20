package app

import (
	"github.com/gin-gonic/gin"
	"yofio/infrastructure/controllers"
	"yofio/infrastructure/database_client"
)

func MapUrls() *gin.Engine {

	db := database_client.GetConnectionPostgres()

	// Provide db variable to controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	ping := router.Group("/")
	{
		ping.GET("/ping", controllers.Ping)
	}

	assigment := router.Group("/")
	{
		assigment.POST("/credit-assignment", controllers.Create)
		assigment.POST("//statistics", controllers.GetStatf)
	}



	return router
}