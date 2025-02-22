package api

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	HTTPClient *http.Client
	APIKey     string
	BaseURL    string
}

func NewClient(apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: false,
			},
		},
		APIKey:  apiKey,
		BaseURL: "https://api.openai.com/v1",
	}
}

func (c *Client) Chat(ctx context.Context, prompt string, stream bool) (string, error) {
	body := fmt.Sprintf(`{
		"model": "%s",
		"messages": [{"role": "user", "content": "%s"}],
		"stream": %t
	}`, model, prompt, stream)

	req, _ := http.NewRequest("POST", c.BaseURL+"/chat/completions", bytes.NewBufferString(body))
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	// 流式处理
	if stream {
		return handleStreamingResponse(req)
	}
	// 普通处理
	resp, err := c.HTTPClient.Do(req)
	// ... 处理普通响应
}

// 流式响应处理
func handleStreamingResponse(req *http.Request) (string, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	var result strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")
			if data == "[DONE]" {
				break
			}

			var chunk struct {
				Choices []struct{ Delta struct{ Content string } }
			}
			if json.Unmarshal([]byte(data), &chunk) == nil {
				fmt.Print(chunk.Choices[0].Delta.Content)
				result.WriteString(chunk.Choices[0].Delta.Content)
			}
		}
	}

	return result.String(), nil
}
