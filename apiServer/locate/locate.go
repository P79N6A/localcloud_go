package locate

import (
	"localcloud_go/util/rabbitmq"
	"os"
	"localcloud_go/util/define"
	"time"
	"strconv"
)

func Locate(name string) string {
	q := rabbitmq.New(os.Getenv(define.RABBITMQ_SERVER))
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}