package platform

import (
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/patrickmn/go-cache"
	"net/url"
	"time"
)

const (
	CacheKeyVerifyTicket         = "web::wx::platform::verifyticket::%s"
	CacheKeyComponentAccessToken = "web::wx::platform::componentaccesstoken::%s"
)

type Client struct {
	http         *httpclient.HttpClient
	ramCache     *cache.Cache
	AppId        string
	AppSecret    string
	EncodeAesKey string
	Token        string
}

func NewClient(appId string, appSecret string, encodeAesKey string, token string) *Client {
	return &Client{
		http:         httpclient.NewHttpClient(),
		ramCache:     cache.New(5*time.Minute, 10*time.Minute),
		AppId:        appId,
		AppSecret:    appSecret,
		EncodeAesKey: encodeAesKey,
		Token:        token,
	}
}

/**
 * 方式一：授权注册页面扫码授权
 * @param $preAuthCode
 * @param $redirectUri
 * @param int $authType
 * @param string $bizAppId
 * @return string
 */
func (c *Client) GetComponentLoginPage(preAuthCode string, redirectUri string, authType int, bizAppId string) string {
	if len(bizAppId) > 0 {
		return fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&biz_appid=%s",
			c.AppId, preAuthCode, url.QueryEscape(redirectUri), bizAppId)
	}
	return fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s",
		c.AppId, preAuthCode, url.QueryEscape(redirectUri), authType)
}

/**
 * 点击移动端链接快速授权
 * 第三方平台方可以生成授权链接，将链接通过移动端直接发给授权管理员，管理员确认后即授权成功。
 * @param $preAuthCode
 * @param $redirectUri
 * @param $authType
 * @param string $bizAppId
 * @return string
 */
func (c *Client) GetBindComponentUrl(preAuthCode string, redirectUri string, authType int, bizAppId string) string {
	if len(bizAppId) > 0 {
		return fmt.Sprintf("https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&biz_appid=%s#wechat_redirect",
			c.AppId, preAuthCode, url.QueryEscape(redirectUri), bizAppId)
	}
	return fmt.Sprintf("https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s#wechat_redirect",
		c.AppId, preAuthCode, url.QueryEscape(redirectUri), authType)
}
