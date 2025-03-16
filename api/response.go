package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func handleStreamResponse(body io.Reader) (chan string, error) {
	ch := make(chan string)
	go func() {
		defer close(ch)
		reader := bufio.NewReader(body)
		thoughtStarted := false

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				ch <- fmt.Sprintf("错误: %v", err)
				return
			}

			if strings.TrimSpace(line) == "" {
				continue
			}

			if strings.HasPrefix(line, "data: ") {
				data := strings.TrimPrefix(line, "data: ")
				data = strings.TrimSpace(data)

				// 处理结束标记
				if data == "[DONE]" {
					break
				}

				var chunk struct {
					Choices []struct {
						Delta struct {
							Content string `json:"content"`
							Thought string `json:"reasoning_content"`
						} `json:"delta"`
					} `json:"choices"`
				}

				if err := json.Unmarshal([]byte(data), &chunk); err != nil {
					ch <- fmt.Sprintf("解析数据块失败: %v", err)
					continue // 继续处理后续数据而不是直接返回
				}

				if len(chunk.Choices) > 0 {
					// 输出思考过程和内容
					if chunk.Choices[0].Delta.Thought != "" {
						if !thoughtStarted {
							ch <- "\n🤔 "
							thoughtStarted = true
						}
						ch <- chunk.Choices[0].Delta.Thought
					}
					if chunk.Choices[0].Delta.Content != "" {
						ch <- chunk.Choices[0].Delta.Content
					}
				}
			}
		}
	}()
	return ch, nil
}

func handleNormalResponse(body io.Reader) (string, error) {
	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("无可用响应")
	}

	return response.Choices[0].Message.Content, nil
}
