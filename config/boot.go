package config

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

var db_config Config

var default_dirs = [...]string{"config.json", "config/config.json"}

func InitDefaultConfig() {
	is_load := false
	for _, path := range default_dirs {
		if gfile.IsFile(path) {
			abs_path := gfile.Abs(path)
			cfg := g.Cfg()
			cfg.SetFileName(path)
			cfg.SetPath(gfile.Dir(abs_path))
			cfg.ToStruct(&db_config)
			is_load = true
			glog.Infof("成功读取日志信息,日志文件路径:%s", abs_path)
			break
		}
	}
	if is_load == false {
		glog.Warningf("读取日志文件失败,请检查以下路径是否存在配置文件!")
	}
}

func Default() *Config {
	return &db_config
}
