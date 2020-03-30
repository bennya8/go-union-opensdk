package mp

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/patrickmn/go-cache"
	"mctea/app/webservices/wx/mp/resp"
	"net/url"
	"time"
)

const (
	CacheKeyAccessToken = "web::wx::mp::accesstoken::%s"
	CacheKeyJsTicket    = "web::wx::mp::jsticket::%s"
)

type Client struct {
	http      *httpclient.HttpClient
	ramCache  *cache.Cache
	AppId     string
	AppSecret string
}

func NewClient() *Client {
	return &Client{
		http:     httpclient.NewHttpClient(),
		ramCache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *Client) GetOauthUrl(redirectUri string, scope string) string {
	u := "https://open.weixin.qq.com/connect/oauth2/authorize?"
	u += "appid=" + c.AppId
	u += "&redirect_uri=" + url.QueryEscape(redirectUri)
	u += "&response_type=code"
	u += "&scope=" + scope
	u += "&state=STATE#wechat_redirect"
	return u
}

func (c *Client) GetSignPackage(url string) (*resp.GetSignPackageRsp, error) {
	var rs resp.GetSignPackageRsp
	//ticket,err := c.GetJsApiTicket()
	//timestamp := time.Now().Unix()
	//nonceStr := helpers.StringRandom(16)
	//rawString := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s")
	//mac := hmac.New(sha1.New, []byte(rawString))
	//_, err := mac.Write([]byte(signStr))
	return &rs, nil
}

func (c *Client) GetAccessToken() (*resp.GetAccessTokenRsp, error) {
	cacheId := fmt.Sprintf(CacheKeyAccessToken, c.AppId)
	var rs resp.GetAccessTokenRsp

	val, exist := c.ramCache.Get(cacheId)
	if !exist {
		rsp, err := c.http.Get("https://api.weixin.qq.com/cgi-bin/token", map[string]interface{}{
			"grant_type": "client_credential",
			"appid":      c.AppId,
			"secret":     c.AppSecret,
		})
		if err != nil {
			return nil, err
		}
		body, err := rsp.ToString()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(body), &rs)
		if err != nil {
			return nil, err
		}
		c.ramCache.Set(cacheId, body, 7200*time.Second)
		return &rs, nil
	}

	err := json.Unmarshal([]byte(val.(string)), &rs)
	if err != nil {
		return nil, err
	}
	return &rs, nil
}

func (c *Client) GetUserAccessToken(code string) (*resp.GetUserAccessTokenRsp, error) {
	var rs resp.GetUserAccessTokenRsp

	rsp, err := c.http.Get("https://api.weixin.qq.com/sns/oauth2/access_token", map[string]interface{}{
		"grant_type": "authorization_code",
		"appid":      c.AppId,
		"secret":     c.AppSecret,
		"code":       code,
	})
	if err != nil {
		return nil, err
	}
	body, err := rsp.ToString()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &rs)
	if err != nil {
		return nil, err
	}
	return &rs, nil
}

func (c *Client) GetUserInfo(userAccessToken string, openId string, lang string) (*resp.GetUserInfoRsp, error) {
	var rs resp.GetUserInfoRsp

	rsp, err := c.http.Get("https://api.weixin.qq.com/sns/userinfo", map[string]interface{}{
		"access_token": userAccessToken,
		"openid":       openId,
		"lang":         lang,
	})
	if err != nil {
		return nil, err
	}
	body, err := rsp.ToString()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &rs)
	if err != nil {
		return nil, err
	}
	return &rs, nil
}

func (c *Client) GetJsApiTicket() (*resp.GetJsApiTicketRsp, error) {
	cacheId := fmt.Sprintf(CacheKeyJsTicket, c.AppId)
	var rs resp.GetJsApiTicketRsp
	val, exist := c.ramCache.Get(cacheId)
	if !exist {
		tokenRsp, err := c.GetAccessToken()
		if err != nil {
			return nil, err
		}
		rsp, err := c.http.Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket", map[string]interface{}{
			"type":         "jsapi",
			"access_token": tokenRsp.AccessToken,
		})
		if err != nil {
			return nil, err
		}
		body, err := rsp.ToString()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(body), &rs)
		if err != nil {
			return nil, err
		}
		c.ramCache.Set(cacheId, body, 7200*time.Second)
		return &rs, nil
	}

	err := json.Unmarshal([]byte(val.(string)), &rs)
	if err != nil {
		return nil, err
	}
	return &rs, nil
}

// 发送模板消息
func (c *Client) MessageTemplateSend() (*resp.MessageTemplateSendResp, error) {
	var rs resp.MessageTemplateSendResp
	rsp, err := c.http.Get("https://api.weixin.qq.com/cgi-bin/message/template/send", map[string]interface{}{
		"access_token": c.GetAccessToken(),
	})
	if err != nil {
		return nil, err
	}
	body, err := rsp.ToString()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &rs)
	if err != nil {
		return nil, err
	}
	return &rs, nil
}
