package openstack

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"golang-demo/common/web"
	"golang-demo/config"
)

const token_key = "X-Subject-Token"

func Auth(r *ghttp.Request) {
	var rest web.Result
	token := r.Session.GetString(token_key)
	if len(token) > 0 {
		rest = web.Result{true, "当前已登录,无需重复登录!", token}
	} else {
		openstack := config.Default().Openstack
		body := `{"auth":{"identity":{"methods":["password"],"password":{"user":{"name":"%s","domain":{"name":"%s"},"password":"%s"}}}}}`
		body = fmt.Sprintf(body, openstack.Username, openstack.Domain, openstack.Password)
		resp, err := ghttp.NewClient().SetContentType("application/json").
			Post(openstack.Url+"/v3/auth/tokens", body)
		if err != nil {
			rest = web.Result{false, err.Error(), nil}
		} else {
			token := resp.Header.Get(token_key)
			if len(token) > 0 {
				r.Session.Set(token_key, token)
			}
			rest = web.Result{true, "登录成功!", token}
		}
	}
	r.Response.WriteJson(rest)
}