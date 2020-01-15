package main

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	_ "golang-demo/boot"
	"golang-demo/config"
)

func main() {
	cfg := config.Default().Server
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	localAddr := fmt.Sprintf("0.0.0.0:%d", cfg.Port)
	server := ghttp.GetServer()

	server.SetAddr(localAddr)
	// 使用驼峰转换
	server.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	//server.EnableAdmin()
	server.EnablePProf()
	server.Start()
	glog.Info("应用启动成功: http://" + addr)
	ghttp.Wait()
}
