package req

/*
ToUserName	开发者微信号
FromUserName	发送方帐号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	消息类型，图片为image
PicUrl	图片链接（由系统生成）
MediaId	图片消息媒体id，可以调用获取临时素材接口拉取数据。
MsgId	消息id，64位整型
*/
type Image struct {
	ToUserName   string `json:"to_user_name"`
	FromUserName string `json:"from_user_name"`
	CreateTime   int64  `json:"create_time"`
	MsgType      string `json:"msg_type"`
	PicUrl       string `json:"pic_url"`
	MediaId      string `json:"media_id"`
	MsgId        int64  `json:"msg_id"`
}

func NewImage() *Image {
	return &Image{
		ToUserName:   "",
		FromUserName: "",
		CreateTime:   0,
		MsgType:      "",
		PicUrl:       "",
		MediaId:      "",
		MsgId:        0,
	}
}
