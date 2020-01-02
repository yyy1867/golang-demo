package boot

import (
	"golang-demo/common/zookeeper"
	"golang-demo/config"
)

func init() {
	// 读取配置文件
	config.InitDefaultConfig()
	// 注册ZK
	zookeeper.InitDefaultConnn()
	// 配置日志
	InitLog()
	// 配置路由
	InitRouter()
}
