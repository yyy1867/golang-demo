package openstack

import (
	"github.com/gogf/gf/net/ghttp"
	"golang-demo/common"
)

const token_key = "X-Subject-Token"

func Auth(r *ghttp.Request) {
	var rest common.Result
	token := r.Session.GetString(token_key)
	if len(token) > 0 {
		rest = common.Result{true, "当前已登录,无需重复登录!", token}
	} else {
		body := `{"auth":{"identity":{"methods":["password"],"password":{"user":{"name":"guxing","domain":{"name":"Default"},"password":"123"}}}}}`
		resp, err := ghttp.NewClient().SetContentType("application/json").
			Post("http://192.168.56.100:5000//v3/auth/tokens", body)
		if err != nil {
			rest = common.Result{false, err.Error(), nil}
		} else {
			token := resp.Header.Get(token_key)
			if len(token) > 0 {
				r.Session.Set(token_key, token)
			}
			rest = common.Result{true, "登录成功!", token}
		}
	}
	r.Response.WriteJson(rest)
}
