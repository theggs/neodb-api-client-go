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

	// 打印原始响应的调试信息
	log.Printf("Raw response: %s\n", resp.String())

	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}

	// 打印解析后的响应调试信息
	responseBody, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format response body: %v", err)
	}
	log.Printf("Parsed response body: %s\n", responseBody)

	return nil
}
