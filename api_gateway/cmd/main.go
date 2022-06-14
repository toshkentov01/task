package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/toshkentov01/task/api_gateway/api/routes"
	"github.com/toshkentov01/task/api_gateway/config"
	"github.com/toshkentov01/task/api_gateway/pkg/middleware"
	"github.com/toshkentov01/task/api_gateway/pkg/utils"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/toshkentov01/task/api_gateway/api/docs" //register swagger
)

var (
	fiberConfig = config.FiberConfig()
	appConfig   = config.Config()
)

// @title TASK API
// @version 0.1
// @description This is an auto-generated API Docs for TASK.
// @termsOfService http://swagger.io/terms/
// @contact.name Sardor Toshkentov
// @contact.email toshkentovsardor.2003@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	app := fiber.New(fiberConfig)

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	middleware.FiberMiddleware(app)

	jwtRoleAuthorizer, err := middleware.NewJWTRoleAuthorizer(appConfig)
	if err != nil {
		log.Fatal("Could not initialize JWT Role Authorizer")
	}

	app.Use(middleware.NewAuthorizer(jwtRoleAuthorizer))
	routes.SwaggerRoute(app)
	routes.CrudRoutes(app)

	// Start server (with or without graceful shutdown).
	if config.Config().Environment == "develop" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
