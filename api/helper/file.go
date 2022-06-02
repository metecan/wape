package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"mime/multipart"
	"path"
	"sync"
	"time"

	"github.com/metecan/wape/api/utils"

	"github.com/beanstalkd/go-beanstalk"
	"github.com/gofiber/fiber/v2"
)

func CheckAndUpload(index int, file *multipart.FileHeader, uploads map[string]string, bucket Bucket, queue Queue, connect *beanstalk.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	// Creating a string for printing indexed file name
	indexedFileName := fmt.Sprintf("%d. %s", index+1, file.Filename)

	// Getting file's header content type
	fileType := file.Header[fiber.HeaderContentType][0]

	// Checking file type
	if fileType == fiber.MIMETextHTML {

		// Opening the file
		buffer, err := file.Open()

		if err != nil {
			uploads[indexedFileName] = "can't uploaded"
		} else {

			// Generating random file name with the original file ext.
			rand.Seed(time.Now().UnixNano())
			generatedFileName := utils.GenerateRandom() + path.Ext(file.Filename)

			// Uploading file to bucket and getting public address
			publicAddr := bucket.Upload(generatedFileName, buffer)

			// Creating the job model and marshalling
			jobModel := make(map[string]string)
			jobModel["path"] = publicAddr
			job, err := json.Marshal(jobModel)
			if err != nil {
				fmt.Println("Error occured: ", err)
			}

			// Putting the job to the queue
			_, err = queue.Put(job, connect)
			if err != nil {
				fmt.Println("Error occured: ", err)
			}

			// if it is uploaded and putted, then adding to uploads map
			uploads[indexedFileName] = "uploaded as " + generatedFileName
		}

		defer buffer.Close()

	} else {

		// If file type is not HTML, then it will be skipped
		uploads[indexedFileName] = "skipped"
	}
}
