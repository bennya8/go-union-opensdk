package mp

import (
	"encoding/json"
	"github.com/bennya8/go-union-opensdk/wx/mp/req"
	"github.com/bennya8/go-union-opensdk/wx/mp/resp"
)

// @link https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#2

// 设置所属行业
func (c *Client) TemplateSetIndustry() {

}

// 获取设置的行业信息
func (c *Client) TemplateGetIndustry() {

}

// 获得模板ID
func (c *Client) TemplateAddTemplate() {

}

// 获取模板列表
func (c *Client) TemplateGetAllTemplate() {

}

// 删除模板
func (c *Client) TemplateDelTemplate() {

}

// 发送模板消息
func (c *Client) TemplateSend(req req.MessageTemplateSendReq) (*resp.MessageTemplateSendResp, error) {
	var rs resp.MessageTemplateSendResp
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var params map[string]interface{}

	paramBytes, _ := json.Marshal(req)
	json.Unmarshal(paramBytes, &params)

	rsp, err := c.http.Post("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+token.AccessToken, params)
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
