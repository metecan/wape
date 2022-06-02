package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

type Queue struct {
	Addr    string
	Port    string
	Network string
}

func InitQueue() *Queue {
	return &Queue{
		Addr:    os.Getenv("QUEUE_ADDR"),
		Port:    os.Getenv("QUEUE_PORT"),
		Network: os.Getenv("QUEUE_NETWORK"),
	}
}

func (q Queue) Connect() (*beanstalk.Conn, error) {

	// Creating a string as full address for connection
	fullAddress := fmt.Sprintf("%s:%s", q.Addr, q.Port)

	// Connecting to the queue (beanstalkd)
	connect, err := beanstalk.Dial(q.Network, fullAddress)

	if err != nil {
		return nil, err
	}

	return connect, nil
}

func (q Queue) Put(body []byte, conn *beanstalk.Conn) (uint64, error) {

	// Putting the job to queue and getting the job id
	id, err := conn.Put(body, 1, 0, 120*time.Second)

	if err != nil {
		return 0, err
	}

	return id, nil
}
