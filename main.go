package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/go-gin-gorm-starter/config"
	"github.com/reynaldineo/go-gin-gorm-starter/models"
	"github.com/reynaldineo/go-gin-gorm-starter/routes"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// custom routes
	routes.UserRoute(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}

// https://gin-gonic.com/docs/quickstart/#getting-started
