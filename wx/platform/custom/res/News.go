package res

type News struct {
	ToUserName   string        `json:"to_user_name"`
	FromUserName string        `json:"from_user_name"`
	CreatedTime  string        `json:"created_time"`
	MsgType      string        `json:"msg_type"`
	ArticleCount int           `json:"article_count"`
	Articles     []interface{} `json:"articles"`
}

func NewNews() *News {
	return &News{
		ToUserName:   "",
		FromUserName: "",
		CreatedTime:  "",
		MsgType:      "news",
		ArticleCount: 0,
		Articles:     nil,
	}
}
