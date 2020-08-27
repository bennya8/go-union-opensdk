package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
	"os"
	"sync"
)

var (
	instance        *cdn.Client
	once            sync.Once
	regionId        = "<regionId>"
	accessKeyID     = "<accessKeyID>"
	accessKeySecret = "<accessKeySecret>"
)

func NewClient() *cdn.Client {
	once.Do(func() {
		regionId = os.Getenv("ALI_SDK_CDN_REGION_ID")
		accessKeyID = os.Getenv("ALI_SDK_CDN_APP_ID")
		accessKeySecret = os.Getenv("ALI_SDK_CDN_APP_SECRET")

		var err error
		instance, err = cdn.NewClientWithAccessKey(regionId, accessKeyID, accessKeySecret)
		if err != nil {
			panic(err)
		}
	})
	return instance
}

func RefreshObjectCaches(client *cdn.Client, objectType string, objectPath string) (response *cdn.RefreshObjectCachesResponse, err error) {

	request := cdn.CreateRefreshObjectCachesRequest()
	request.Scheme = "https"

	request.ObjectPath = objectPath
	request.ObjectType = objectType

	return client.RefreshObjectCaches(request)
}
