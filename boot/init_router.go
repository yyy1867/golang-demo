package boot

import (
	"github.com/gogf/gf/net/ghttp"
	openstack "golang-demo/ctrl/api"
)

func init() {
	server := ghttp.GetServer()
	initOpenstackRouter(server)
}

func initOpenstackRouter(server *ghttp.Server) {
	server.Group("open", func(group *ghttp.RouterGroup) {
		group.GET("auth", openstack.Auth)
	})
}
