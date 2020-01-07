package redis

import (
	"errors"
	"github.com/go-redis/redis/v7"
	"github.com/gogf/gf/os/glog"
	"golang-demo/config"
)

var default_client redis.UniversalClient

func Default() redis.UniversalClient {
	if default_client == nil {
		cfg := config.Default().Redis
		if len(cfg.Addrs) == 1 {
			default_client = redis.NewClient(&redis.Options{
				Addr:     cfg.Addrs[0],
				Password: cfg.Password,
			})
		} else if len(cfg.Addrs) > 1 {
			default_client = redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:    cfg.Addrs,
				Password: cfg.Password,
			})
		}
		if default_client == nil {
			panic(errors.New("初始化redis连接失败,请检查redis配置项!"))
		}
		cmd := default_client.Exists("test")
		if cmd.Err() != nil {
			glog.Errorf("redis连接失败:%s", cmd.Err().Error())
			panic(cmd.Err())
		}
		glog.Info("redis连接初始化成功!")
	}

	return default_client
}
