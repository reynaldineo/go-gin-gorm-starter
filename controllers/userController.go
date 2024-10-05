package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/go-gin-gorm-starter/config"
	"github.com/reynaldineo/go-gin-gorm-starter/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	config.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
