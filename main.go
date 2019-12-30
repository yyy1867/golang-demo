package main

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	_ "golang-demo/boot"
	"golang-demo/common"
)

func main() {
	addr := fmt.Sprintf("%s:%d", common.DbConfig.Server.Host, common.DbConfig.Server.Port)
	localAddr := fmt.Sprintf("0.0.0.0:%d", common.DbConfig.Server.Port)
	server := ghttp.GetServer()
	server.SetAddr(localAddr)
	// 使用驼峰转换
	server.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	//server.EnableAdmin()
	server.Start()
	glog.Info("应用启动成功: http://" + addr)
	ghttp.Wait()
}
