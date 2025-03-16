package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DeepSeekClient struct {
	HTTPClient *http.Client
	APIKey     string
	BaseURL    string
}

func (c *DeepSeekClient) Name() string {
	return "deepseek"
}

func NewDeepSeekClient(baseURL string, apiKey string) *DeepSeekClient {
	return &DeepSeekClient{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: false,
			},
		},
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
}

func (c *DeepSeekClient) CreateChatCompletion(prompt string, model string, stream bool) (chan string, error) {
	reqBody := map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"stream": stream,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %v", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/chat/completions", bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("网络请求失败: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("读取错误响应失败: %v", err)
		}
		return nil, fmt.Errorf("API 错误 (状态码 %d): %s", resp.StatusCode, string(body))
	}

	if stream {
		return handleStreamResponse(resp.Body)
	}

	content, err := handleNormalResponse(resp.Body)
	resp.Body.Close()

	ch := make(chan string, 1)
	ch <- content
	close(ch)
	return ch, err
}
