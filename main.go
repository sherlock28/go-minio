package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/sherlock28/go-minio/api/routes"
	"github.com/sherlock28/go-minio/api/utils"
)

func main() {

	app := fiber.New()

	routes.PublicRoutes(app)

	utils.StartServer(app)
}
