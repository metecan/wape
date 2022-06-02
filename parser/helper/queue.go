package helper

import (
	"fmt"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

type Queue struct {
	Addr    string
	Port    string
	Network string
}

type JobModel struct {
	Path string `json:"path"`
}

func InitQueue() *Queue {
	return &Queue{
		Addr:    "127.0.0.1",
		Port:    "11300",
		Network: "tcp",
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

func (q Queue) Reserve(conn *beanstalk.Conn) (uint64, []byte, error) {

	// Getting the jobs from queue
	id, body, err := conn.Reserve(50 * time.Second)
	defer conn.Delete(id)

	if err != nil {
		return 0, nil, err
	}

	return id, body, nil
}
