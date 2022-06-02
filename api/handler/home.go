package handler

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	// Returning basic JSON to prove the server is alive
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":    "Wape API is alive",
		"status": "ok",
	})
}
