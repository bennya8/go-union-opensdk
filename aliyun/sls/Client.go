package sls

import (
	sdk "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/bennya8/go-union-opensdk/utils"
	"os"
	"sync"
	"time"
)

var (
	endpoint        = "<endpoint>"
	accessKeyID     = "<accessKeyId>"
	accessKeySecret = "<accessKeySecret>"
	instance        *sdk.Client
	once            sync.Once
)

func NewClient() *sdk.Client {
	once.Do(func() {
		accessKeyID = os.Getenv("ALI_SDK_SLS_APP_ID")
		accessKeySecret = os.Getenv("ALI_SDK_SLS_APP_SECRET")
		endpoint = os.Getenv("ALI_SDK_SLS_END_POINT")

		instance = &sdk.Client{
			Endpoint:        endpoint,
			AccessKeyID:     accessKeyID,
			AccessKeySecret: accessKeySecret,
		}
	})
	return instance
}

func PutLogLog(project string, store string, params map[string]string) error {
	client := NewClient()

	var content []*sdk.LogContent
	for k, v := range params {
		content = append(content, &sdk.LogContent{
			Key:   utils.ConvertString(k),
			Value: utils.ConvertString(v),
		})
	}
	log := &sdk.Log{
		Time:     utils.ConvertUint32(uint32(time.Now().Unix())),
		Contents: content,
	}

	logGroup := &sdk.LogGroup{
		Topic: utils.ConvertString(""),
		//Source: proto.String("10.230.201.117"),
		Logs: []*sdk.Log{log},
	}

	return client.PutLogs(project, store, logGroup)
}
