package req

type RequestMsg map[string]interface{}

type IRequestMsg interface {
	GetRequest() RequestMsg
}
