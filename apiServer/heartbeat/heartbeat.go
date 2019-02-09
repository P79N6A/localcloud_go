package heartbeat

import (
	"localcloud_go/util/rabbitmq"
	"os"
	"time"
	"localcloud_go/util/define"
	"sync"
	"strconv"
)

var dataServers = make(map[string]time.Time)
/**
 dataServers需要mutex的保护，以防多个goroutine并发读写map造成错误。
 Go语言的map可以支持多个goroutine同时读，但不能支持多个goroutine同时写或同时既读又写
 所以这里使用一个互斥所mutex保护map的并发读写
 无论读写都只允许一个goroutine操作map。另一个更有效率的方法是使用读写锁RWMutex，因为读写锁可以允许多个goroutine同时读
*/
var mutex sync.Mutex

/**
 创建消息队列绑定apiServers
 */
func ListenHeartBeat() {
	q := rabbitmq.New(os.Getenv(define.RABBITMQ_SERVER))
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

/**
 每隔5s扫描一遍dataServers，并清除10s没有收到心跳消息的数据服务节点
 */
func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

/**
  遍历dataServers并返回当前所有的数据服务节点
 */
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s := range dataServers {
		ds = append(ds, s)
	}
	return ds
}

func StartHeartBeat() {
	q := rabbitmq.New(os.Getenv(define.RABBITMQ_SERVER))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv(define.LISTEN_ADDRESS))
		time.Sleep(5 * time.Second)
	}
}