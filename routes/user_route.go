package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/go-gin-gorm-starter/controller"
	"github.com/reynaldineo/go-gin-gorm-starter/middleware"
	"github.com/reynaldineo/go-gin-gorm-starter/service"
)

func UserRoute(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("", userController.Register)
		routes.POST("/login", userController.Login)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.GetMe)
	}
}
