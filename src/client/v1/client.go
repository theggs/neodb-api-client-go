package v1

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

// NeoDBClient 用于与 neodb.social API 交互的客户端
type NeoDBClient struct {
	apiBaseURL string
	client     *resty.Client
}

// NewNeoDBClient 创建一个新的 NeoDBClient 实例
func NewNeoDBClient(apiBaseURL string) *NeoDBClient {
	return &NeoDBClient{
		apiBaseURL: apiBaseURL,
		client:     resty.New(),
	}
}

// executeRequest 执行 API 请求并解析响应
func (c *NeoDBClient) executeRequest(endpoint string, result interface{}) error {
	request := c.client.R().
		SetHeader("Accept", "application/json")

	// TODO: 添加授权头部，使用 access_token
	// request.SetAuthToken(c.accessToken)

	// 打印调试信息
	url := fmt.Sprintf("%s%s", c.apiBaseURL, endpoint)
	log.Printf("Sending request to URL: %s", url)

	resp, err := request.Get(url)
	if err != nil {
		return fmt.Errorf("failed to execute request: %v", err)
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("failed to execute request: status code %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return nil
}
