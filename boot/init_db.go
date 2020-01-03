package boot

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/glog"
	"golang-demo/config"
)

func InitDatabase() {
	mysql := config.Default().Mysql
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Database)
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Type:     "mysql",
				Debug:    true,
				LinkInfo: url,
			},
		},
	})
	glog.Info("默认数据库配置成功!")
}
