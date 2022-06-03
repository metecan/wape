package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func ReserveAndInsert(queue Queue, connect *beanstalk.Conn, db *sql.DB, wg *sync.WaitGroup) {

	defer wg.Done()

	// Getting the body of the job from the queue
	_, body, err := queue.Reserve(connect)
	if err != nil {
		fmt.Println(err)
	}

	// Creating JobModel model and binding the body to the struct
	var jsonBody JobModel
	json.Unmarshal(body, &jsonBody)

	// getting all anchors from the file
	anchors, err := Parse(jsonBody.Path)
	if err != nil {
		fmt.Println(err)
	}

	// Spliting the path to getting file name
	fileName := strings.Split(jsonBody.Path, "/")[9]

	// Spliting the file name to getting the keyword
	fileKeyword := strings.Split(fileName, ".")[0]

	// Getting the current timestamp
	currentEpoch := time.Now().Unix()

	// Created a variable to check how many titles were added to DB
	insertedCount := 0

	// Inserting every title (anchor) to the DB
	for _, anchor := range anchors {
		err := db.QueryRow("INSERT INTO titles (keyword,title,epoch) VALUES($1,$2,$3)", fileKeyword, anchor, currentEpoch).Err()
		if err != nil {
			fmt.Println("Title can't inserted to database")
		}
		insertedCount++
	}

	// Checking insertedCount
	if insertedCount == len(anchors) {
		fmt.Printf("All titles in %s file inserted to db\n", fileName)
	}
}
