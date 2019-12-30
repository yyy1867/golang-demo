package boot

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"golang-demo/ctrl/api/zookeeper"
	"time"
)

func init() {
	client, _, err := zk.Connect([]string{"192.168.56.102:2181"}, time.Second*5)
	if err != nil {
		fmt.Println("连接zk失败:" + err.Error())
	}
	zookeeper.ZkClient = client
}
