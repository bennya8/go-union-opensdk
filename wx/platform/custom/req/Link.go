package req

/*
ToUserName	接收方微信号
FromUserName	发送方微信号，若为普通用户，则是一个OpenID
CreateTime	消息创建时间
MsgType	消息类型，链接为link
Title	消息标题
Description	消息描述
Url	消息链接
MsgId	消息id，64位整型
*/
type Link struct {
	ToUserName   string `json:"to_user_name"`
	FromUserName string `json:"from_user_name"`
	CreateTime   int64  `json:"create_time"`
	MsgType      string `json:"msg_type"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Url          string `json:"url"`
	MsgId        int64  `json:"msg_id"`
}

func NewLink() *Link {
	return &Link{
		ToUserName:   "",
		FromUserName: "",
		CreateTime:   0,
		MsgType:      "",
		Title:        "",
		Description:  "",
		Url:          "",
		MsgId:        0,
	}
}
