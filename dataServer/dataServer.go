package main

import (
	"localcloud_go/dataServer/heartbeat"
	"localcloud_go/dataServer/locate"
	"net/http"
	"localcloud_go/dataServer/objects"
	"git.elenet.me/MaxQ/log"
	"os"
	"localcloud_go/util/define"
)

func main()  {
	os.Setenv(define.RABBITMQ_SERVER, "amqp://test:test@127.0.0.1:5672")
	os.Setenv(define.ES_SERVER, "127.0.0.1:9200")
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv(define.LISTEN_ADDRESS), nil))

}
