package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/sherlock28/go-minio/api/routes"
	"github.com/sherlock28/go-minio/api/utils"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	routes.PublicRoutes(app)

	utils.StartServer(app)
}
