package sendcloud

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var instance *Client
var once sync.Once

type Client struct {
	mailUser string
	mailKey  string
	smsUser  string
	smsKey   string
	http     *http.Client
}

func NewClient(smsUser string, smsKey string, mailUser string, mailKey string) *Client {
	once.Do(func() {
		instance = &Client{}
		instance.smsUser = smsUser
		instance.smsKey = smsKey
		instance.mailUser = mailUser
		instance.mailKey = mailKey
		instance.http = &http.Client{}
	})
	return instance
}

func (c *Client) SendSms(tplId string, mobile string, vars map[string]string) (string, error) {
	uri := "http://www.sendcloud.net/smsapi/send"

	params := url.Values{
		"phone":      {mobile},
		"smsUser":    {c.smsUser},
		"templateId": {tplId},
	}

	varsBytes, _ := json.Marshal(vars)
	signature := c.signature(params, string(varsBytes))
	params.Add("vars", string(varsBytes))
	params.Add("signature", signature)

	return c.request("POST", uri, params)
}

func (c *Client) SendMail(from string, fromName string, tos []string, subject string, html string) (string, error) {
	uri := "http://sendcloud.sohu.com/webapi/mail.send.json"
	//不同于登录SendCloud站点的帐号，您需要登录后台创建发信子帐号，使用子帐号和密码才可以进行邮件的发送。
	params := url.Values{
		"api_user": {c.mailUser},
		"api_key":  {c.mailKey},
		"from":     {from},
		"fromname": {fromName},
		"to":       {strings.Join(tos, ";")},
		"subject":  {subject},
		"html":     {html},
	}
	return c.request("POST", uri, params)
}

func (c *Client) signature(params url.Values, vars string) string {
	paramStr := params.Encode()
	if vars != "" {
		paramStr += "&vars=" + vars
	}
	paramStr = c.smsKey + "&" + paramStr + "&" + c.smsKey
	m := md5.New()
	m.Write([]byte(paramStr))
	return hex.EncodeToString(m.Sum(nil))
}

func (c *Client) request(method, url string, params url.Values) (string, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(params.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	response := map[string]interface{}{}
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return "", err
	}

	if _, ok := response["errors"]; ok {
		return "", errors.New(fmt.Sprintf("%s", response["errors"]))
	}

	return string(responseBytes), nil

}
