package boot

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/samuel/go-zookeeper/zk"
	"golang-demo/common"
	"time"
)

func init() {
	client, _, err := zk.Connect(common.DbConfig.Zookeeper.Addrs, time.Second*1000)
	client.SetLogger(glog.DefaultLogger())
	if err != nil {
		glog.Error("连接zk失败:%s", err.Error())
	}
	registered(client)
	zookeeper.ZkClient = client
	glog.Info("zk连接初始化成功!")
}

func registered(client *zk.Conn) {
	server := common.DbConfig.Server
	addr := fmt.Sprintf("%s:%d", server.Host, server.Port)
	path := "/bocloud/services/godemo/" + addr
	paths := gstr.Split(path, "/")
	for i := range paths {
		if i > 0 {
			tpath := gstr.Join(paths[0:i+1], "/")
			on, _, err := client.Exists(tpath)
			if err != nil {
				glog.Errorf("判断节点是否存在失败[%s]:%s", tpath, err.Error())
				return
			}
			if on == false {
				var flag int32 = zk.FlagSequence & zk.FlagEphemeral
				if i == len(paths)-1 {
					flag = zk.FlagEphemeral
				}
				node, err := client.Create(tpath, []byte(paths[i]), flag, zk.WorldACL(zk.PermAll))
				if err != nil {
					glog.Errorf("创建zk节点失败[%s]:%s", tpath, err.Error())
					return
				}
				glog.Infof("创建zk节点成功:[%s]:%s", tpath, node)
			}
		}
	}
}
