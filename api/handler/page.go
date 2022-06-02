package handler

import (
	"sync"

	"github.com/metecan/wape/api/helper"

	"github.com/gofiber/fiber/v2"
)

var wg sync.WaitGroup

func Page(c *fiber.Ctx) error {

	// initialize bucket
	bucket := helper.InitBucket()

	// initialize queue
	queue := helper.InitQueue()

	// Connecting to queue (beanstalkd)
	connect, err := queue.Connect()
	if err != nil {
		panic(err)
	}

	defer connect.Close()

	// Creating map for controls
	uploads := make(map[string]string)

	// Getting multipart-form data from body
	form, _ := c.MultipartForm()

	// Checking form if it is empty
	if form != nil {
		// Getting files which sent with the "files" prefix
		files := form.File["file"]

		// Checking files count
		if len(files) > 0 {
			// Creating loop for each file
			for index, file := range files {
				// Adding the proccess to waitgroup
				wg.Add(1)

				// Starting file checking & uploadin at the same time with goroutines
				go helper.CheckAndUpload(index, file, uploads, *bucket, *queue, connect, &wg)
			}
		} else {
			// If there is no file, then returning an error
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error",
				"msg":    "Please upload files with 'files' prefix",
			})
		}

		// Waiting for goroutines
		wg.Wait()

		// Checking files that skipped
		skippedCount := 0

		for _, upload := range uploads {
			if upload == "skipped" || upload == "can't uploaded" {
				skippedCount++
			}
		}

		// If skipped files are equal to all files length, then returning an information-like error
		if len(files) == skippedCount {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error",
				"msg":    "any file uploaded to bucket.",
			})
		}

		// If everything is ok, then returning an information.
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  "ok",
			"msg":     "HTML files uploaded to bucket",
			"uploads": uploads,
		})

	} else {
		// If form body is empty, then returning an error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"msg":    "Please upload files",
		})
	}
}
