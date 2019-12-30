package boot

import (
	"github.com/gogf/gf/os/glog"
	"github.com/samuel/go-zookeeper/zk"
	"golang-demo/common"
	"golang-demo/ctrl/api/zookeeper"
	"time"
)

func init() {
	client, _, err := zk.Connect(common.DbConfig.Zookeeper.Addrs, time.Second*1000)
	client.SetLogger(glog.DefaultLogger())
	if err != nil {
		glog.Error("连接zk失败:%s", err.Error())
	}
	zookeeper.ZkClient = client
	glog.Info("zk连接初始化成功!")
}
