package handler

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestPage(t *testing.T) {

	app := fiber.New()

	// http.Request
	req := httptest.NewRequest("GET", "/page", nil)

	// http.Response
	resp, _ := app.Test(req)

	// Do something with results:
	if resp.StatusCode == fiber.StatusMethodNotAllowed {
		fmt.Println(resp.StatusCode)
	}
}
