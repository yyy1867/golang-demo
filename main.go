package main

import (
	"github.com/gogf/gf/net/ghttp"
	_ "golang-demo/boot"
)

func main() {
	server := ghttp.GetServer()
	server.SetAddr("127.0.0.1:8888")
	server.EnableAdmin()
	server.Run()
}
