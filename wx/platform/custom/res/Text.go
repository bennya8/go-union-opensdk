package res

type Text struct {
	ToUserName   string `json:"to_user_name"`
	FromUserName string `json:"from_user_name"`
	CreateTime   string `json:"create_time"`
	MsgType      string `json:"msg_type"`
	Content      string `json:"content"`
}

func NewText() *Text {
	return &Text{
		ToUserName:   "",
		FromUserName: "",
		CreateTime:   "",
		MsgType:      "text",
		Content:      "",
	}
}
