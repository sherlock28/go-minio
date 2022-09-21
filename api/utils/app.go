package utils

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	a.Listen(":" + port)
	fmt.Println("Server on port " + port)
}
