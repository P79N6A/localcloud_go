package locate

import (
	"os"
	"localcloud_go/util/rabbitmq"
	"localcloud_go/util/define"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(os.Getenv(define.RABBITMQ_SERVER))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		if Locate(os.Getenv(define.STORAGE_ROOT) + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv(define.LISTEN_ADDRESS))
		}
	}
}
