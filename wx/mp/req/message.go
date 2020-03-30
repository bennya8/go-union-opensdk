package req

type MessageTemplateSendReq struct {
	Touser      string `json:"touser"`
	TemplateID  string `json:"template_id"`
	URL         string `json:"url"`
	Miniprogram struct {
		Appid    string `json:"appid"`
		Pagepath string `json:"pagepath"`
	} `json:"miniprogram"`
	Data map[string]MessageTemplateSendDataReq `json:"data"`
}

type MessageTemplateSendDataReq struct {
	Value string `json:"value"`
	Color string `json:"color"`
}
