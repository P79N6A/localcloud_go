package heartbeat

import (
	"time"
	"math/rand"
)

/**
 当前所有数据服务节点中随机选出一个节点并返回
 */
func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	return ds[rand.Intn(n)]
}
