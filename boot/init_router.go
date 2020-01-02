package boot

import (
	"github.com/gogf/gf/net/ghttp"
	"golang-demo/app/openstack"
	"golang-demo/app/zookeeper"
)

func init() {
	server := ghttp.GetServer()
	initOpenstackRouter(server)
	initZookeeperRouter(server)
}

func initOpenstackRouter(server *ghttp.Server) {
	server.Group("open", func(group *ghttp.RouterGroup) {
		group.GET("auth", openstack.Auth)
	})
}

func initZookeeperRouter(server *ghttp.Server) {
	zk := new(zookeeper.ZkController)
	server.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	server.BindController("/zk/{.method}", zk)
}
