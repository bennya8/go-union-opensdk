package req

import "time"

type Event struct {
	ToUserName   string    `json:"to_user_name"`
	FromUserName string    `json:"from_user_name"`
	CreateTime   string `json:"create_time"`
	MsgType      string    `json:"msg_type"`
	Event        string    `json:"event"`
	EventKey     string    `json:"event_key"`
	Ticket       string    `json:"ticket"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	SuccTime time.Time `json:"succ_time"`
	FailTime time.Time `json:"fail_time"`
}
