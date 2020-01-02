package boot

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"time"
)

func InitLog() {
	logger := glog.DefaultLogger()
	logger.SetLevel(glog.LEVEL_ALL)
	logger.SetPath("log")
	logger.SetStdoutPrint(true)
	bindRouteLogger()
	glog.Info("日志配置完成!")
}

func bindRouteLogger() {
	server := ghttp.GetServer()
	server.BindMiddlewareDefault(func(r *ghttp.Request) {
		// 开始时间
		start := time.Now()
		r.Middleware.Next()
		end := time.Now()
		latency := end.Sub(start)
		path := r.Request.URL.Path
		clientIP := r.Host
		method := r.Method
		status := r.Response.Status
		glog.Infof("%s %-3d %-13v %s %s", method, status, latency, clientIP, path)
	})
}
