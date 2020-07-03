package boot

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang-demo/app/deploy"
	"golang-demo/app/openstack"
	"golang-demo/app/zookeeper"
)

func InitRouter() {
	server := ghttp.GetServer()
	initOpenstackRouter(server)
	initZookeeperRouter(server)
	initMyDeployRouter(server)
	initPrometheus(server)
}

func initOpenstackRouter(server *ghttp.Server) {
	server.Group("open", func(group *ghttp.RouterGroup) {
		group.GET("auth", openstack.Auth)
		group.GET("server-list", openstack.ServerList)
	})
}

func initZookeeperRouter(server *ghttp.Server) {
	zkctrl := new(zookeeper.ZkController)
	server.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	server.BindController("/zk/{.method}", zkctrl)
}

func initMyDeployRouter(server *ghttp.Server) {
	obj := &deploy.MyDeploy{}
	server.BindObject("/deploy/{.method}", obj)
}

func initPrometheus(server *ghttp.Server) {
	handler := promhttp.Handler()
	server.BindHandler("/metrics", func(r *ghttp.Request) {
		handler.ServeHTTP(r.Response.Writer, r.Request)
	})
}
