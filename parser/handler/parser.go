package handler

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/metecan/wape/parser/db"
	"github.com/metecan/wape/parser/helper"

	"github.com/gofiber/fiber/v2"
)

var wg sync.WaitGroup

func Parser(c *fiber.Ctx) error {
	queue := helper.InitQueue()

	// Connecting the queue (beanstalkd)
	connect, err := queue.Connect()
	if err != nil {
		fmt.Println(err)
	}

	defer connect.Close()

	// Getting jobs' status
	jobs, err := connect.Stats()
	if err != nil {
		fmt.Println(err)
	}

	// Getting the jobs that are currently ready
	currentJobs, err := strconv.Atoi(jobs["current-jobs-ready"])
	if err != nil {
		fmt.Println(err)
	}

	// Connecting to the DB
	db, err := db.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Checking count of jobs
	if currentJobs > 0 {
		fmt.Println("Process Starting...")

		// Running process for every job in the queue with goroutines
		for i := 0; i < currentJobs; i++ {
			fmt.Printf("%d. Job in the process\n", i+1)

			wg.Add(1)
			go helper.ReserveAndInsert(*queue, connect, db, &wg)
		}
	} else {
		// if there is no job, then returning error-like output
		return fiber.NewError(fiber.StatusBadRequest, "there is no job in the queue")
	}

	// Waiting for all goroutines
	wg.Wait()

	return nil
}
