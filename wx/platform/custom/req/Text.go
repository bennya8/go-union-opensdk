package req

/*
ToUserName	开发者微信号
FromUserName	发送方帐号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	消息类型，文本为text
Content	文本消息内容
MsgId	消息id，64位整型
*/
type Text struct {
	ToUserName   string `json:"to_user_name"`
	FromUserName string `json:"from_user_name"`
	CreateTime   int64  `json:"create_time"`
	MsgType      string `json:"msg_type"`
	Content      string `json:"content"`
	MsgId        int64  `json:"msg_id"`
}

func NewText() *Text {
	return &Text{
		ToUserName:   "",
		FromUserName: "",
		CreateTime:   0,
		MsgType:      "",
		Content:      "",
		MsgId:        0,
	}
}
