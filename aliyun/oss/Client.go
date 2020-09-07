package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bennya8/go-union-opensdk/utils"
	"math/rand"
	"os"
	"strconv"
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

func GetObjectDomain() string {
	if len(domain) > 0 {
		return domain
	}
	return endpoint
}

func GetObjectDir() string {
	return utils.TimeFormatToNowString("YYYY") + "/" + utils.TimeFormatToNowString("MM") + "/" + utils.TimeFormatToNowString("DD") + "/"
}
func GetObjectKey() string {
	now := utils.TimeFormatToNowString()
	return utils.CryptMD5(now + strconv.Itoa(rand.Intn(1000000)) + strconv.Itoa(rand.Intn(1000000)))
}
