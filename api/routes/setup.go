package routes

import (
	"github.com/metecan/wape/api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Home)
	app.Post("/page", handler.Page)
}
