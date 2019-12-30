package main

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	_ "golang-demo/boot"
)

func main() {
	addr := "127.0.0.1:8888"
	server := ghttp.GetServer()
	server.SetAddr(addr)
	// 使用驼峰转换
	server.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	//server.EnableAdmin()
	server.Start()
	glog.Info("应用启动成功: http://" + addr)
	ghttp.Wait()
}
