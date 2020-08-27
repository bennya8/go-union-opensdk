package nls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	regionId        = "<regionId>"
	accessKeyId     = "<accessKeyId>"
	accessKeySecret = "<accessKeySecret>"
	once            sync.Once
	instance        *sdk.Client
)

func NewClient() (*sdk.Client, error) {
	once.Do(func() {
		accessKeyId = os.Getenv("ALI_SDK_NLS_APP_ID")
		accessKeySecret = os.Getenv("ALI_SDK_NLS_APP_SECRET")

		var err error
		instance, err = sdk.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
		if err != nil {
			panic(err)
		}
	})
	return instance, nil
}

func processGETRequest(token string, text string, audioSaveFile string, format string, sampleRate int, voice string) {
	/**
	 * 设置HTTPS GET请求
	 * 1.使用HTTPS协议
	 * 2.语音识别服务域名：nls-gateway.cn-shanghai.aliyuncs.com
	 * 3.语音识别接口请求路径：/stream/v1/tts
	 * 4.设置必须请求参数：appkey、token、text、format、sample_rate
	 * 5.设置可选请求参数：voice、volume、speech_rate、pitch_rate
	 */
	var url = "https://nls-gateway.cn-shanghai.aliyuncs.com/stream/v1/tts"
	url = url + "?appkey=" + accessKeyId
	url = url + "&token=" + token
	url = url + "&text=" + text
	url = url + "&format=" + format
	url = url + "&sample_rate=" + strconv.Itoa(sampleRate)
	// voice 发音人，可选，默认是xiaoyun
	url = url + "&voice=" + voice
	// volume 音量，范围是0~100，可选，默认50
	// url = url + "&volume=" + strconv.Itoa(50)
	// speech_rate 语速，范围是-500~500，可选，默认是0
	// url = url + "&speech_rate=" + strconv.Itoa(0)
	// pitch_rate 语调，范围是-500~500，可选，默认是0
	// url = url + "&pitch_rate=" + strconv.Itoa(0)
	fmt.Println(url)
	/**
	 * 发送HTTPS GET请求，处理服务端的响应
	 */
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("The GET request failed!")
		panic(err)
	}
	defer response.Body.Close()
	contentType := response.Header.Get("Content-Type")
	body, _ := ioutil.ReadAll(response.Body)
	if "audio/mpeg" == contentType {
		file, _ := os.Create(audioSaveFile)
		defer file.Close()
		file.Write([]byte(body))
		fmt.Println("The GET request succeed!")
	} else {
		// ContentType 为 null 或者为 "application/json"
		statusCode := response.StatusCode
		fmt.Println("The HTTP statusCode: " + strconv.Itoa(statusCode))
		fmt.Println("The GET request failed: " + string(body))
	}
}

func processPOSTRequest(token string, text string, audioSaveFile string, format string, sampleRate int) {
	/**
	 * 设置HTTPS POST请求
	 * 1.使用HTTPS协议
	 * 2.语音合成服务域名：nls-gateway.cn-shanghai.aliyuncs.com
	 * 3.语音合成接口请求路径：/stream/v1/tts
	 * 4.设置必须请求参数：appkey、token、text、format、sample_rate
	 * 5.设置可选请求参数：voice、volume、speech_rate、pitch_rate
	 */
	var url string = "https://nls-gateway.cn-shanghai.aliyuncs.com/stream/v1/tts"
	bodyContent := make(map[string]interface{})
	bodyContent["appkey"] = accessKeyId
	bodyContent["token"] = token

	bodyContent["text"] = text
	bodyContent["format"] = format
	bodyContent["sample_rate"] = sampleRate
	// voice 发音人，可选，默认是xiaoyun
	// bodyContent["voice"] = "xiaoyun"
	// volume 音量，范围是0~100，可选，默认50
	// bodyContent["volume"] = 50
	// speech_rate 语速，范围是-500~500，可选，默认是0
	// bodyContent["speech_rate"] = 0
	// pitch_rate 语调，范围是-500~500，可选，默认是0
	bodyContent["pitch_rate"] = 100
	bodyJson, err := json.Marshal(bodyContent)
	if err != nil {
		panic(nil)
	}
	fmt.Println(string(bodyJson))
	/**
	 * 发送HTTPS POST请求，处理服务端的响应
	 */
	response, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte(bodyJson)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	contentType := response.Header.Get("Content-Type")
	body, _ := ioutil.ReadAll(response.Body)
	if "audio/mpeg" == contentType {
		file, _ := os.Create(audioSaveFile)
		defer file.Close()
		file.Write([]byte(body))
		fmt.Println("The POST request succeed!")
	} else {
		// ContentType 为 null 或者为 "application/json"
		statusCode := response.StatusCode
		fmt.Println("The HTTP statusCode: " + strconv.Itoa(statusCode))
		fmt.Println("The POST request failed: " + string(body))
	}
}

func main() {
	var token string = "5753f3b70aac4abf820e68a307b10722"
	var text string = "如果您准备好了就点击确定：如果，您准备好了，就点击确定"
	var textUrlEncode = text
	textUrlEncode = url.QueryEscape(textUrlEncode)
	textUrlEncode = strings.Replace(textUrlEncode, "+", "%20", -1)
	textUrlEncode = strings.Replace(textUrlEncode, "*", "%2A", -1)
	textUrlEncode = strings.Replace(textUrlEncode, "%7E", "~", -1)
	fmt.Println(textUrlEncode)
	var audioSaveFile = "demo_t296_q10_by_xiaobei.mp3"
	var format = "mp3"
	var sampleRate int = 16000
	processGETRequest(token, textUrlEncode, audioSaveFile, format, sampleRate, "xiaobei")
	// processPOSTRequest(appkey, token, text, audioSaveFile, format, sampleRate)

	audioSaveFile = "demo_t296_q10_by_xiaomei.mp3"
	processGETRequest(token, textUrlEncode, audioSaveFile, format, sampleRate, "xiaomei")
}

func GetToken() {
	client, err := NewClient()
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Domain = "nls-meta.cn-shanghai.aliyuncs.com"
	request.Version = "2018-05-18"
	request.PathPattern = "/pop/2018-05-18/tokens"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpStatus())
	fmt.Print(response.GetHttpContentString())
}
