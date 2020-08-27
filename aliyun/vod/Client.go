package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	regionId        = "<regionId>"
	accessKeyId     = "<accessKeyId>"
	accessKeySecret = "<accessKeySecret>"
	instance        *vod.Client
	once            sync.Once
)

type TemplateGroupId string

func NewClient() (*vod.Client, error) {
	once.Do(func() {
		// get from env
		regionId = os.Getenv("ALI_SDK_VOD_REGION_ID")
		accessKeyId = os.Getenv("ALI_SDK_VOD_APP_ID")
		accessKeySecret = os.Getenv("ALI_SDK_VOD_APP_SECRET")

		// 创建授权对象
		credential := &credentials.AccessKeyCredential{
			accessKeyId,
			accessKeySecret,
		}
		// 自定义config
		config := sdk.NewConfig()
		config.AutoRetry = true     // 失败是否自动重试
		config.MaxRetryTime = 3     // 最大重试次数
		config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒

		// 创建vodClient实例
		var err error
		instance, err = vod.NewClientWithOptions(regionId, config, credential)
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance, nil
}

func UploadMediaByURL(client *vod.Client, ossInUrls []string, metas string, templateGroupId TemplateGroupId) (*vod.UploadMediaByURLResponse, error) {
	request := vod.CreateUploadMediaByURLRequest()
	request.TemplateGroupId = string(templateGroupId)
	request.UploadURLs = strings.Join(ossInUrls, ",")
	request.AcceptFormat = "JSON"
	if len(metas) > 0 {
		request.UploadMetadatas = metas
	}
	return client.UploadMediaByURL(request)
}

func GetURLUploadInfos(client *vod.Client, ossInUrls []string) (*vod.GetURLUploadInfosResponse, error) {
	request := vod.CreateGetURLUploadInfosRequest()
	request.UploadURLs = strings.Join(ossInUrls, ",")
	request.AcceptFormat = "JSON"
	return client.GetURLUploadInfos(request)
}

func MyGetPlayInfo(client *vod.Client, videoId string) (*vod.GetPlayInfoResponse, error) {
	request := vod.CreateGetPlayInfoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.GetPlayInfo(request)
}

func MyGetPlayAuth(client *vod.Client, videoId string) (*vod.GetVideoPlayAuthResponse, error) {
	request := vod.CreateGetVideoPlayAuthRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	request.AuthInfoTimeout = "3000"
	return client.GetVideoPlayAuth(request)
}
