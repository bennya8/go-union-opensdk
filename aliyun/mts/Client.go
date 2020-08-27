package mts

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/mts"
	"log"
	"net/url"
	"os"
	"sync"
)

var (
	regionId        = "<regionId>"
	bucket          = "<bucket>"
	location        = "<location>"
	accessKeyId     = "<accessKeyId>"
	accessKeySecret = "<accessKeySecret>"
	instance        *mts.Client
	once            sync.Once
)

type PipelineId string
type TemplateId string

func NewClient() (*mts.Client, error) {
	once.Do(func() {
		// get from env
		regionId = os.Getenv("ALI_SDK_MTS_REGION_ID")
		accessKeyId = os.Getenv("ALI_SDK_MTS_APP_ID")
		accessKeySecret = os.Getenv("ALI_SDK_MTS_APP_SECRET")
		bucket = os.Getenv("ALI_SDK_MTS_BUCKET")
		location = os.Getenv("ALI_SDK_MTS_LOCATION")

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

		var err error
		instance, err = mts.NewClientWithOptions(regionId, config, credential)
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance, nil
}

// 转码接口 - 提交转码作业
func SubmitJobs(client *mts.Client, inObjId string, outObjId string, templateId TemplateId, pipelineId PipelineId) (*mts.SubmitJobsResponse, error) {
	inputJson, _ := json.Marshal(map[string]string{
		"Bucket":   bucket,
		"Location": location,
		"Object":   url.QueryEscape(inObjId),
	})

	outputJson, _ := json.Marshal([]map[string]string{
		{
			"OutputObject": url.QueryEscape(outObjId),
			"TemplateId":   string(templateId),
		},
	})

	request := mts.CreateSubmitJobsRequest()
	request.PipelineId = string(pipelineId)
	request.Input = string(inputJson)
	request.OutputBucket = bucket
	request.OutputLocation = location
	request.Outputs = string(outputJson)

	response, err := client.SubmitJobs(request)
	if err != nil {
		return nil, err
	}
	return response, err
}

// 转码接口 - 取消转码作业
func CancelJob() {}

// 转码接口 - 查询转码作业
func QueryJobList() {

}

// 转码接口 - 列出转码作业
func ListJob() {

}

func AddTemplate() {

}

func UpdateTemplate() {

}

func QueryTemplateList() {

}

func SearchTemplate() {

}

func DeleteTemplate() {

}
