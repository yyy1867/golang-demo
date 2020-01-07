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
	query := db.Table("tbl_cop2life_column_info")
	total, _ := query.Count()
	err := query.Limit((param.Page-1)*param.Rows, param.Rows).Structs(&columns)
	if err != nil {
		rst = web.Result{false, "查询失败:" + err.Error(), nil}
		g.Log().Warningf("查询失败:%s", err.Error())
	} else {
		rst = web.Result{true, "查询成功!", web.GridBean{
			Page:  param.Page,
			Pages: len(columns),
			Total: total,
			Rows:  columns,
		}}
	}
	r.Response.WriteJson(rst)
}

func (m *MyDeploy) Show(r *ghttp.Request) {
	test := r.Session.GetString("test")
	r.Response.Write(test)
}

func (m *MyDeploy) Set(r *ghttp.Request) {
	pmap := r.GetMap()
	if pmap["test"] == nil {
		pmap["test"] = "默认值"
	}
	for k, v := range pmap {
		r.Session.Set(k, v)
	}
	r.Response.WriteJsonExit(web.Result{true, "写入会话成功!", pmap})
}
