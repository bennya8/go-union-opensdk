package req

/*
ToUserName	开发者微信号
FromUserName	发送方帐号（一个OpenID）
CreateTime	消息创建时间 （整型）
MsgType	消息类型，地理位置为location
Location_X	地理位置维度
Location_Y	地理位置经度
Scale	地图缩放大小
Label	地理位置信息
MsgId	消息id，64位整型
*/
type Location struct {
	ToUserName   string  `json:"to_user_name"`
	FromUserName string  `json:"from_user_name"`
	CreateTime   int64   `json:"create_time"`
	MsgType      string  `json:"msg_type"`
	LocationX    float64 `json:"location_x"`
	LocationY    float64 `json:"location_y"`
	Scale        int     `json:"scale"`
	Label        string  `json:"label"`
	MsgId        int64   `json:"msg_id"`
}

func NewLocation() *Location {
	return &Location{
		ToUserName:   "",
		FromUserName: "",
		CreateTime:   0,
		MsgType:      "location",
		LocationX:    0,
		LocationY:    0,
		Scale:        0,
		Label:        "",
		MsgId:        0,
	}
}
