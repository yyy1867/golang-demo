package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"golang-demo/common/redis"
	"time"
)

func InitSessionConfig() {
	server := ghttp.GetServer()
	storage := redis.NewStorageRedis(redis.Default(), "golang.demo.")
	server.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute * 3,
		"SessionStorage": storage,
	})
	g.Log("session管理器配置完毕!")
}
