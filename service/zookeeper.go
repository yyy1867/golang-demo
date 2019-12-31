package service

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/samuel/go-zookeeper/zk"
	"golang-demo/common"
)

const (
	CopUser        BoCloudService = "copuser"
	CopParamCenter BoCloudService = "copparamcenter"
	CopCycle       BoCloudService = "copcycle"
	CopTask        BoCloudService = "coptask"
	CopDeploy      BoCloudService = "copdeploy"
)

const (
	CREATE  BoCloudMethod = "POST"
	MODIFY  BoCloudMethod = "POST"
	REMOVE  BoCloudMethod = "DELETE"
	LIST    BoCloudMethod = "POST"
	LOAD    BoCloudMethod = "GET"
	BASIC   BoCloudMethod = "GET"
	SUBOBJ  BoCloudMethod = "GET"
	OPERATE BoCloudMethod = "POST"
	OTHERS  BoCloudMethod = "POST"
)

const BoCloudZKRootPath string = "/bocloud/services/"

type BoCloudMethod string

type BoCloudService string

type ServiceFactory struct {
	Client *zk.Conn
}

func (s *ServiceFactory) ListServices(service BoCloudService) []string {
	path := fmt.Sprintf("%s%s", BoCloudZKRootPath, service)
	children, _, err := s.Client.Children(path)
	if err != nil {
		glog.Warningf("获取ZK路径[%s]失败:%s", path, err.Error())
	}
	return children
}

func (s *ServiceFactory) Request(service BoCloudService, uri string, method BoCloudMethod,
	header map[string]string, param map[string]interface{}) common.BsmResult {
	ips := s.ListServices(service)
	if len(ips) == 0 {
		return common.BsmResult{false, "远程服务未找到!", nil, "", ""}
	}
	i := gtime.Timestamp() % int64(len(ips))
	hurl := "http://" + ips[i] + uri
	client := ghttp.NewClient()
	if len(header) > 0 {
		client.SetHeaderMap(header)
	}
	mtd := fmt.Sprintf("%s", method)
	if "GET" == mtd {
		urlparam := ghttp.BuildParams(param)
		hurl = hurl + "?" + urlparam
		param = nil
	}
	content := client.RequestContent(mtd, hurl, param)
	var result common.BsmResult
	err := gjson.DecodeTo(content, &result)
	if err != nil {
		result = common.BsmResult{true, "格式转换失败:" + err.Error(), content, "", ""}
	}
	return result
}
