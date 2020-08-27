package sms

import (
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"os"
	"sync"
)

var (
	instance        *dysmsapi.Client
	once            sync.Once
	regionId        = "<regionId>"
	accessKeyId     = "<accessKeyId>"
	accessKeySecret = "<accessKeySecret>"
)

func NewClient() (*dysmsapi.Client, error) {
	once.Do(func() {
		// get from env
		regionId = os.Getenv("ALI_SDK_SMS_REGION_ID")
		accessKeyId = os.Getenv("ALI_SDK_SMS_APP_ID")
		accessKeySecret = os.Getenv("ALI_SDK_SMS_APP_SECRET")

		var err error
		instance, err = dysmsapi.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance, nil
}

func SendSms(client *dysmsapi.Client, signName string, phones string, tplId string, tplParams map[string]string) (response *dysmsapi.SendSmsResponse, err error) {
	if len(phones) <= 0 {
		return nil, errors.New("phones is required")
	}
	if len(tplId) <= 0 {
		return nil, errors.New("tpl is required")
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phones
	request.TemplateCode = tplId
	request.SignName = signName

	tplParamsJson, _ := json.Marshal(tplParams)
	if len(tplParamsJson) > 0 {
		request.TemplateParam = string(tplParamsJson)
	}

	// 上行短信扩展码，无特殊需要此字段的用户请忽略此字段。
	//request.SmsUpExtendCode
	// 外部流水扩展字段。
	//request.OutId

	return client.SendSms(request)
}
