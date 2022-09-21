package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sherlock28/go-minio/api/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// upload files
	route.Post("/upload", controllers.UploadFile)

	route.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"api": "Upload file with GO API in MinIO",
		})
	})
}
