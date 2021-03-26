package zcy_sdk_go

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	GET 	= "GET"
	POST	= "POST"
	PUT		= "PUT"
	DELETE	= "DELETE"
)

type AssetCloudRequest struct {
	// 完整的请求路径 https://platform.assetcloud.org.cn/dev-api/+请求路径
	Url string
	// 平台获取的 key
	Key string
	// 平台获取的 secret
	Secret string
	// POST和PUT请求的body
	Body string
	// HTTP method
	HttpMethod string
}

type AssetCloudResponse struct {
	// 状态码
	Code int				`json:"code"`
	// 是否成功
	Success bool			`json:"success"`
	// 承载数据
	Data json.RawMessage	`json:"data"`
	// 返回消息
	Msg string				`json:"msg"`
}

func Send(cloudRequest *AssetCloudRequest) *AssetCloudResponse {
	client := &http.Client{}

	bodyBytes := []byte(cloudRequest.Body)

	// url拼接时间戳+加签
	url := handleRequest(cloudRequest.Url, cloudRequest.Secret)

	request, e := http.NewRequest(cloudRequest.HttpMethod, url, bytes.NewBuffer(bodyBytes))
	if e != nil {
		os.Exit(1)
	}
	request.Header.Add("key", cloudRequest.Key)
	// POST、PUT采用json传输
	if cloudRequest.HttpMethod == "POST" || cloudRequest.HttpMethod == "PUT" {
		request.Header.Add("Content-Type", "application/json")
	}

	// 发起请求
	response, e := client.Do(request)
	if e != nil {
		os.Exit(1)
	}
	defer response.Body.Close()
	// 读取请求返回结果
	respBytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		os.Exit(1)
	}

	resp := AssetCloudResponse{}
	e = json.Unmarshal(respBytes, &resp)
	if e != nil {
		os.Exit(1)
	}
	return &resp
}

func handleRequest(url, secret string) string {
	param := ""
	timestamp := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	// 拼接时间戳
	if strings.Contains(url, "?") {
		url += "&timestamp=" + timestamp
	} else {
		url += "?timestamp=" + timestamp
	}
	if strings.Contains(url, "?") {
		param = url[strings.Index(url, "?")+1:]
	}
	// 加签
	sign := hmacUrl(secret, param)
	url += "&sign=" + sign
	return url
}

// 拼接hmac
func hmacUrl(secret, url string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(url))
	sha := hex.EncodeToString(hash.Sum(nil))
	return string([]byte(sha))
}
