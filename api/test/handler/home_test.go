package handler

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHome(t *testing.T) {

	app := fiber.New()

	// http.Request
	req := httptest.NewRequest("GET", "/", nil)

	// http.Response
	resp, _ := app.Test(req)

	// Do something with results:
	if resp.StatusCode == fiber.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		if string(body) == `{"msg":"Wape API is alive","status":"ok"}` {
			fmt.Println(string(body))
		}
	}
}
