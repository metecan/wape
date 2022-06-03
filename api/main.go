package main

import (
	"log"

	"github.com/metecan/wape/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
