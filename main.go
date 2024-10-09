package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/reynaldineo/go-gin-gorm-starter/cmd"
	"github.com/reynaldineo/go-gin-gorm-starter/config"
	"github.com/reynaldineo/go-gin-gorm-starter/controller"
	"github.com/reynaldineo/go-gin-gorm-starter/middleware"
	"github.com/reynaldineo/go-gin-gorm-starter/repository"
	"github.com/reynaldineo/go-gin-gorm-starter/routes"
	"github.com/reynaldineo/go-gin-gorm-starter/service"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		cmd.Commands(db)
		return
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		//* === Dependecy Injection Implementation ===

		//* == Repository ==
		userRepository repository.UserRepository = repository.NewUserRepository(db)

		//* == Service ==
		userService service.UserService = service.NewUserService(userRepository, jwtService)

		//* == Controller ==
		userController controller.UserController = controller.NewUserController(userService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	//* == routes ==
	routes.UserRoute(server, userController, jwtService)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "development" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}

}
