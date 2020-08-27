package dtplus

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/bennya8/go-union-opensdk/utils"
	"github.com/imroc/req"
	"log"
	http2 "net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	regionId        = ""
	accessKeyID     = ""
	accessKeySecret = ""
	instance        *sdk.Client
	once            sync.Once
)

func NewClient() *sdk.Client {
	once.Do(func() {
		regionId = os.Getenv("ALI_SDK_DTP_REGION_ID")
		accessKeyID = os.Getenv("ALI_SDK_DTP_APP_ID")
		accessKeySecret = os.Getenv("ALI_SDK_DTP_APP_SECRET")

		client, err := sdk.NewClientWithAccessKey(regionId, accessKeyID, accessKeySecret)
		if err != nil {
			panic(err)
		}
		instance = client
	})
	return instance
}

func FaceVerifyByUrl(image string, image2 string) {
	regionId = os.Getenv("ALS_DTP_REGION_ID")
	accessKeyID = os.Getenv("ALI_DTP_APP_ID")
	accessKeySecret = os.Getenv("ALI_DTP_APP_SECRET")

	http := req.New()
	fmt.Println(http)

	param := req.Param{
		"type":        0,
		"image_url_1": image,
		"image_url_2": image2,
	}
	content, _ := json.Marshal(param)
	fmt.Println(string(content))

	var contentMd5 string

	contentMd5 = utils.CryptMD5Base64(string(content))
	fmt.Println(contentMd5)

	accept := "application/json"
	contentType := "application/json"
	urlPath := "/face/verify"

	loc, _ := time.LoadLocation("GMT")
	date := time.Now().In(loc).Format(http2.TimeFormat)

	sign := signature(accessKeySecret, "POST", accept, contentMd5, contentType, date, urlPath)
	fmt.Println(sign)

	header := req.Header{
		"accept":        accept,
		"content-type":  contentType,
		"date":          date,
		"authorization": fmt.Sprintf("Dataplus %s:%s", accessKeyID, sign),
	}
	fmt.Println(header)

	rsp, err := http.Post("https://dtplus-cn-shanghai.data.aliyuncs.com"+urlPath, header, req.BodyJSON(param))

	log.Printf("%+v", rsp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp.ToString())

	//client := NewClient()
	//
	//request := requests.NewCommonRequest()
	//request.Method = "POST"
	//request.Domain = "dtplus-cn-shanghai.data.aliyuncs.com/face/verify"
	//
	//
	//
	//response, err := client.ProcessCommonRequest(request)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Print(response.GetHttpStatus())
	//fmt.Print(response.GetHttpContentString())

}

func FaceVerifyByUrl2(image string, image2 string) {
	client := NewClient()

	param := map[string]string{
		"type":        strconv.Itoa(0),
		"image_url_1": image,
		"image_url_2": image2,
	}
	content, _ := json.Marshal(param)

	contentMd5 := utils.CryptMD5(string(content))

	accept := "application/json"
	contentType := "application/json"
	urlPath := "/face/attribute"

	loc, _ := time.LoadLocation("GMT")
	date := time.Now().In(loc).Format(http2.TimeFormat)

	sign := signature(accessKeySecret, "POST", accept, contentMd5, contentType, date, urlPath)

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Domain = "dtplus-cn-shanghai.data.aliyuncs.com"
	//request.PathPattern = "/logstores"
	//request.PathParams["logstoreName"] = store
	//request.Headers["x-log-apiversion"] = "0.6.0"
	//request.Headers["Host"] =

	request.Headers["Authorization"] = fmt.Sprintf("Dataplus %s:%s", accessKeyID, sign)
	request.Headers["Content-type"] = contentType
	request.Headers["Accept"] = accept
	request.Headers["Date"] = date

	request.FormParams = param

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpStatus())
	fmt.Print(response.GetHttpContentString())

}

// @todo https://help.aliyun.com/document_detail/67818.html?spm=a2c4g.11186623.6.560.2e692f86HupVQV
// @todo 接入说明
func signature(appSecret string, method string, accept string, contentMd5 string, contentType string, date string, path string) string {
	signStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", method, accept, contentMd5, contentType, date, path)

	mac := hmac.New(sha1.New, []byte(appSecret))
	_, err := mac.Write([]byte(signStr))
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
