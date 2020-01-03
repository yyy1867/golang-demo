package deploy

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"golang-demo/common/web"
)

type MyDeploy struct {
}

func (m *MyDeploy) List(r *ghttp.Request) {
	var rst web.Result
	var param web.PageParam
	r.GetQueryStruct(&param)
	web.DefaultPageParam(&param)
	db := g.DB()
	var columns []ColumnInfo
	err := db.Table("tbl_cop2life_column_info").Limit(param.Page, param.Rows).Structs(&columns)
	if err != nil {
		rst = web.Result{false, "查询失败:" + err.Error(), nil}
		g.Log().Warningf("查询失败:%s", err.Error())
	} else {
		rst = web.Result{true, "查询成功!", columns}
	}
	r.Response.WriteJson(rst)
}
