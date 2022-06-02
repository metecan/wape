package main

import (
	"log"
	"os"

	"github.com/metecan/wape/api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
