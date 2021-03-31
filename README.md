# Go SDK的用法说明
* `url`: 完整的请求路径 https://platform.assetcloud.org.cn/dev-api/+请求路径
* `key`、`secret`: 平台获取的 key 和 secret

### 安装
---
    go get github.com/assetcloud/zcy-sdk-go
### 使用
Get调用方法：
```go
    import (
    	asset "github.com/assetcloud/zcy-sdk-go"
    )

    request := &asset.AssetCloudRequest{
        Url:        "",
        Key:        "",
        Secret:     "",
        Body:       "",
        HttpMethod: asset.GET,
    }
    response := asset.Send(request)
    respBytes, _ := json.Marshal(response)
    println(string(respBytes))
```
Post调用方法：
```go
    import (
    	"encoding/json"
    	asset "github.com/assetcloud/zcy-sdk-go"
    )

    bodyMap := make(map[string]interface{})
    bodyMap["key1"] = "val1"
    bodyBytes, _ := json.Marshal(bodyMap)
    postRequest := &asset.AssetCloudRequest{
        Url:        "",
        Key:        "",
        Secret:     "",
        Body:       string(bodyBytes),
        HttpMethod: asset.POST,
    }
    postResponse := asset.Send(postRequest)
    postRespBytes, _ := json.Marshal(postResponse)
    println(string(postRespBytes))
```
DELETE调用方法：
```go
    import (
        asset "github.com/assetcloud/zcy-sdk-go"
    )

    request := &asset.AssetCloudRequest{
        Url:        "",
        Key:        "",
        Secret:     "",
        Body:       "",
        HttpMethod: asset.DELETE,
    }
    response := asset.Send(request)
    respBytes, _ := json.Marshal(response)
    println(string(respBytes))
```
PUT调用方法：
```go
    import (
    	"encoding/json"
    	asset "github.com/assetcloud/zcy-sdk-go"
    )

    bodyMap := make(map[string]interface{})
    bodyMap["key1"] = "val1"
    bodyBytes, _ := json.Marshal(bodyMap)
    postRequest := &asset.AssetCloudRequest{
        Url:        "",
        Key:        "",
        Secret:     "",
        Body:       string(bodyBytes),
        HttpMethod: asset.PUT,
    }
    postResponse := asset.Send(postRequest)
    postRespBytes, _ := json.Marshal(postResponse)
    println(string(postRespBytes))
```

返回结果为AssetCloudResponse：
| 字段 | 类型 | 说明 |
| ---- | ---- | ---- |
| code | int | 状态码 |
| success | bool | 是否成功 |
| data | T | 承载数据 |
| msg | string | 返回消息 | 