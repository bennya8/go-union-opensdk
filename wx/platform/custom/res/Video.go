package res

type Video struct {
	ToUserName   string `json:"to_user_name"`
	FromUserName string `json:"from_user_name"`
	CreateTime   string `json:"create_time"`
	MsgType      string `json:"msg_type"`
	MediaId      string `json:"media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

func NewVideo() *Video {
	return &Video{
		ToUserName:   "",
		FromUserName: "",
		CreateTime:   "",
		MsgType:      "video",
		MediaId:      "",
		Title:        "",
		Description:  "",
	}
}
