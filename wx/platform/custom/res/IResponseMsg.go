package res

type ResponseMsg map[string]interface{}

type IResponseMsg interface {
	GetResponse() ResponseMsg
}
