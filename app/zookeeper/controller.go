package zookeeper

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
	"github.com/samuel/go-zookeeper/zk"
	"golang-demo/common/web"
	"golang-demo/common/zookeeper"
)

type ZkController struct {
	gmvc.Controller
	client         *zk.Conn
	serviceFactory *ServiceFactory
}

func (c *ZkController) Init(r *ghttp.Request) {
	c.Controller.Init(r)
	c.client = zookeeper.DefaultConn()
	c.serviceFactory = &ServiceFactory{Client: zookeeper.DefaultConn()}
}

func (z *ZkController) That() {
	z.Response.Writef("%p", z)
}

func (z *ZkController) UserList() {
	param := map[string]interface{}{
		"rows": 2,
		"page": 2,
	}
	request := z.serviceFactory.Request(CopUser, "/user/list", LOAD, nil, param)
	z.Response.WriteJson(request)
}

func (z *ZkController) List() {
	var rst web.Result
	children, stat, err := z.client.Children("/")
	if err != nil {
		rst = web.Result{false, "查询失败:" + err.Error(), nil}
	} else {
		rst = web.Result{true, "查询成功!", g.Map{
			"stat":     stat,
			"children": children,
		}}
	}
	z.Response.WriteJson(rst)
}

func (z *ZkController) Tree() {
	rmap := map[string]interface{}{}
	zktree("/", &rmap, z)
	z.Response.WriteJson(web.Result{true, "查询成功!", rmap})
}

func zktree(path string, parent *map[string]interface{}, z *ZkController) {
	children, stat, err := z.client.Children(path)
	childs := make([]map[string]interface{}, len(children))
	*parent = g.Map{
		"path":   path,
		"stat":   stat,
		"error":  err,
		"childs": childs,
	}
	if err == nil && len(childs) > 0 {
		if "/" != path {
			path += "/"
		}
		for i := range childs {
			zktree(path+children[i], &childs[i], z)
		}
	}
}
