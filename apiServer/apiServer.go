package main

import (
	"localcloud_go/apiServer/heartbeat"
	"net/http"
	"localcloud_go/apiServer/objects"
	"localcloud_go/apiServer/locate"
	"git.elenet.me/MaxQ/log"
	"os"
	"localcloud_go/util/define"
	"localcloud_go/apiServer/versions"
)

func main()  {
	os.Setenv(define.RABBITMQ_SERVER, "amqp://test:test@127.0.0.1:5672")
	os.Setenv(define.ES_SERVER, "127.0.0.1:9200")
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv(define.LISTEN_ADDRESS), nil))
}


