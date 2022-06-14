package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/toshkentov01/task/api_gateway/api/controllers"
)

// CrudRoutes ...
func CrudRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET METHOD:
	route.Get("/post/list/", controllers.ListPosts)
	route.Get("/post/get/:post_id/", controllers.GetPost)

	// Routes for PUT METHOD:
	route.Put("/post/update/", controllers.UpdatePost)

	// Routes for DELETE METHOD:
	route.Delete("/post/delete/:post_id/", controllers.DeletePost)
}
