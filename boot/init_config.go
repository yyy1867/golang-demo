package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"golang-demo/common"
)

func init() {
	cfg := g.Cfg()
	cfg.SetFileName("config.json")
	cfg.ToStruct(&common.DbConfig)
	glog.Info("读取配置完毕!")
}
