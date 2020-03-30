package req

/*
ToUserName	开发者微信号
FromUserName	发送方帐号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	小视频为shortvideo
MediaId	视频消息媒体id，可以调用获取临时素材接口拉取数据。
ThumbMediaId	视频消息缩略图的媒体id，可以调用获取临时素材接口拉取数据。
MsgId	消息id，64位整型
*/
type ShortVideo struct {
	ToUserName   string `json:"to_user_name"`
	FromUserName string `json:"from_user_name"`
	CreateTime   int64  `json:"create_time"`
	MsgType      string `json:"msg_type"`
	MediaId      string `json:"media_id"`
	ThumbMediaId string `json:"thumb_media_id"`
	MsgId        int64  `json:"msg_id"`
}

func NewShortVideo() *ShortVideo {
	return &ShortVideo{
		ToUserName:   "",
		FromUserName: "",
		CreateTime:   0,
		MsgType:      "",
		MediaId:      "",
		ThumbMediaId: "",
		MsgId:        0,
	}
}
