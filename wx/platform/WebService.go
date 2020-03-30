package platform

import (
	"github.com/bennya8/go-union-opensdk/wx/platform/custom/req"
	"github.com/bennya8/go-union-opensdk/wx/platform/custom/res"
)

type DecodeMsg struct {
}

type DecodeOauthMsg struct {
	InfoType string `json:"info_type"`
}

type IWxPlatformWebService interface {

	/***********************************
	 * Message Request
	 **********************************/

	MsgTextRequest(text *req.Text) res.IResponseMsg
	MsgImageRequest(image *req.Image) res.IResponseMsg
	MsgVoiceRequest(voice req.Voice) res.IResponseMsg
	MsgVideoRequest(video req.ShortVideo) res.IResponseMsg
	MsgShortVideoRequest(video req.ShortVideo) res.IResponseMsg
	MsgLocationRequest(location req.Location) res.IResponseMsg
	MsgLinkRequest(link req.Link) res.IResponseMsg
	MsgEventRequest(event req.Event) res.IResponseMsg

	/***********************************
	 * Callback Handler
	 **********************************/

	MsgNotifyProcess(appId string, data []byte)
	MsgNotifyProcessOnError(errorCode int, msg string)

	OauthNotifyProcess()
}
