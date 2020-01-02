package zookeeper

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/samuel/go-zookeeper/zk"
	"golang-demo/config"
	"runtime"
	"time"
)

var default_conn *zk.Conn

func InitDefaultConnn() {
	cfg := config.Default().Zookeeper
	default_conn, events, err := zk.Connect(cfg.Addrs, 3*time.Second,
		zk.WithLogger(glog.DefaultLogger()))
	if err != nil {
		glog.Error("连接zk失败:%s", err.Error())
	}
	go eventlistener(events, default_conn)
}

func DefaultConn() *zk.Conn {
	return default_conn
}

func eventlistener(events <-chan zk.Event, client *zk.Conn) {
	cfg := config.Default().Zookeeper
	count := 0
	for e := range events {
		etype := int32(e.Type)
		estate := int32(e.State)
		errstr := ""
		if e.Err != nil {
			errstr = e.Err.Error()
		}
		if e.Type == zk.EventSession && e.State == zk.StateConnecting {
			count++
			if count >= cfg.MaxConnCount {
				glog.Error("zk连接次数[%d],未能成功连接,程序即将关闭!", count)
				client.Close()
				ghttp.GetServer().Shutdown()
				runtime.Goexit()
			}
			continue
		} else if e.Type == zk.EventSession && e.State == zk.StateHasSession {
			count = 0
			registered(client)
		}
		glog.Infof("消息类型:[%d],State:[%d],路径:[%s],异常:[%s],地址:[%s]", etype, estate, e.Path, errstr, e.Server)
	}
}

func registered(client *zk.Conn) {
	server := config.Default().Server
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
	glog.Info("zk节点注册成功!")
}
