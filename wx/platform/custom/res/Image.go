package res

type Image struct {
	ToUserName   string `json:"to_user_name"`
	FromUserName string `json:"from_user_name"`
	CreateTime   string `json:"create_time"`
	MsgType      string `json:"msg_type"`
	MediaId      string `json:"media_id"`
}

func NewImage() *Image {
	return &Image{
		ToUserName:   "",
		FromUserName: "",
		CreateTime:   "",
		MsgType:      "image",
		MediaId:      "",
	}
}
