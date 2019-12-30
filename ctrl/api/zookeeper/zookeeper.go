package zookeeper

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/samuel/go-zookeeper/zk"
	"golang-demo/common"
)

var ZkClient *zk.Conn

type ZkController struct {
	gmvc.Controller
}

func (z *ZkController) List() {
	var rst common.Result
	children, stat, err := ZkClient.Children("/")
	if err != nil {
		rst = common.Result{false, "查询失败:" + err.Error(), nil}
	} else {
		rst = common.Result{true, "查询成功!", g.Map{
			"stat":     stat,
			"children": children,
		}}
	}
	z.Response.WriteJson(rst)
}

func (z *ZkController) Tree() {
	rmap := map[string]interface{}{}
	zktree("/", &rmap)
	z.Response.WriteJson(common.Result{true, "查询成功!", rmap})
}

func zktree(path string, parent *map[string]interface{}) {
	children, stat, err := ZkClient.Children(path)
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
			zktree(path+children[i], &childs[i])
		}
	}
}
