package heartbeat

import (
	"localcloud_go/util/rabbitmq"
	"os"
	"localcloud_go/util/define"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv(define.RABBITMQ_SERVER))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv(define.LISTEN_ADDRESS))
		time.Sleep(5 * time.Second)
	}
}
