package req

/*
ToUserName	开发者微信号
FromUserName	发送方帐号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	语音为voice
MediaId	语音消息媒体id，可以调用获取临时素材接口拉取数据。
Format	语音格式，如amr，speex等
MsgId	消息id，64位整型
*/
type Voice struct {
	ToUserName   string `json:"to_user_name"`
	FromUserName string `json:"from_user_name"`
	CreateTime   int64  `json:"create_time"`
	MsgType      string `json:"msg_type"`
	MediaId      string `json:"media_id"`
	Format       string `json:"format"`
	MsgId        int64  `json:"msg_id"`
}

func NewVoice() *Voice {
	return &Voice{
		ToUserName:   "",
		FromUserName: "",
		CreateTime:   0,
		MsgType:      "",
		MediaId:      "",
		Format:       "",
		MsgId:        0,
	}
}
