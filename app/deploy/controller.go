package deploy

import (
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
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
		rst = web.Result{true, "查询成功!", warpGridBean(gconv.Interfaces(columns), total, param)}
	}
	r.Response.WriteJson(rst)
}

func (m *MyDeploy) Show(r *ghttp.Request) {
	test := r.Session.GetString("test")
	r.Response.WriteJson(web.Result{true, "读取成功!", test})
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

func (m *MyDeploy) Err(r *ghttp.Request) {
	rst := g.Map{
		"text":  gerror.New("文本错误!").Error(),
		"stack": gerror.Stack(gerror.New("堆栈错误!")),
		"warp":  gerror.Wrap(gerror.New("原始错误!"), "包装->").Error(),
		"cause": gerror.Cause(gerror.Wrap(gerror.New("原始错误!"), "包装->")).Error(),
	}
	r.Response.WriteJson(web.Result{true, "成功!", rst})
}

func (m *MyDeploy) Glist(r *ghttp.Request) {
	list := glist.New(false)
	list.PushBack("测试元素1")
	list.PushBack("测试元素2")
	list.PushBack("测试元素3")
	r.Response.WriteJson(web.Result{true, "list操作测试!", list})
}

// 包装为分页对象的方法
func warpGridBean(data []interface{}, total int, page web.PageParam) web.GridBean {
	return web.GridBean{
		Page:  page.Page,
		Pages: len(data),
		Total: total,
		Rows:  data,
	}
}
