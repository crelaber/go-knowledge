package basic

import (
	"fmt"
	"go-knowledge/wechat/officialaccount/context"
	"go-knowledge/wechat/util"
)

var (
	// 获取微信服务器IP地址
	// 文档：https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
	getCallbackIPURL  = "https://api.weixin.qq.com/cgi-bin/getcallbackip"
	getAPIDomainIPURL = "https://api.weixin.qq.com/cgi-bin/get_api_domain_ip"
	// 清理接口调用次数
	clearQuotaURL = "https://api.weixin.qq.com/cgi-bin/clear_quota"
)

type Basic struct {
	*context.Context
}

func NewBasic(context *context.Context) *Basic {
	basic := new(Basic)
	basic.Context = context
	return basic
}

// IPListRes 获取微信的ip服务器地址
type IPListRes struct {
	util.CommonErr
	IPList []string `json:"ip_list"`
}

func (basic *Basic) GetCallbackIP() ([]string, error) {

	ak, err := basic.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?access_token=%s", getCallbackIPURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}

	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetCallbackIP")
	return ipListRes.IPList, err
}

// GetAPIDomainIP 获取微信api接口
func (basic *Basic) GetAPIDomainIP() ([]string, error) {
	ak, err := basic.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getAPIDomainIPURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}

	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetAPIDomainIP")
	return ipListRes.IPList, err
}

// ClearQuota 清理接口调用次数
func (basic *Basic) ClearQuota() error {
	ak, err := basic.GetAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s?access_token=%s", clearQuotaURL, ak)
	data, err := util.PostJSON(url, map[string]string{
		"appid": basic.AppID,
	})
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(data, "ClearQuota")
}
