package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"sync"
)

var (
	endpoint        = "<endpoint>"
	accessKeyID     = "<accessKeyId>"
	accessKeySecret = "<accessKeySecret>"
	domain          = "<domain>"
	instance        *oss.Client
	once            sync.Once
)

func NewClient() *oss.Client {
	once.Do(func() {
		accessKeyID = os.Getenv("ALI_SDK_OSS_APP_ID")
		accessKeySecret = os.Getenv("ALI_SDK_OSS_APP_SECRET")
		endpoint = os.Getenv("ALI_SDK_OSS_END_POINT")
		domain = os.Getenv("ALI_SDK_OSS_CUSTOM_DOMAIN")

		client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
		if err != nil {
			panic(err)
		}
		instance = client
	})
	return instance
}

