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

func (q Queue) Put(body []byte, conn *beanstalk.Conn) (uint64, error) {

	// Putting the job to queue and getting the job id
	id, err := conn.Put(body, 1, 0, 120*time.Second)

	if err != nil {
		return 0, err
	}

	return id, nil
}
