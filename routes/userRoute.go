package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/go-gin-gorm-starter/controllers"
)

func UserRoute(r *gin.Engine) {

	r.GET("/users", controllers.GetUsers)
}
