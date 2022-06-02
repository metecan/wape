package main

import (
	"log"

	"github.com/metecan/wape/parser/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {

	app := fiber.New()

	app.Get("/", handler.Parser)

	log.Fatalln(app.Listen(":8001"))

}
