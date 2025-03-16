package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	interactiveCmd.Flags().StringVarP(&model, "model", "m", "", "模型名称")
	interactiveCmd.Flags().BoolVarP(&stream, "stream", "s", true, "启用流式输出")
	interactiveCmd.Flags().StringVarP(&provider, "provider", "p", "", "指定供应商 (volcengine, deepseek)") // 新增 provider 参数
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "进入交互式对话模式",
	Run: func(cmd *cobra.Command, args []string) {
		initProvider()
		startREPL()
	},
}

/*
进入交互式对话循环
*/
func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("🤖 [%s](%s)\n", model, currentProvider.Name())

	for {
		fmt.Print(">>> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			break
		}

		handleInput(input)
	}
}

/*
处理用户输入
*/
func handleInput(prompt string) {
	stream, err := currentProvider.CreateChatCompletion(prompt, model, true)
	if err != nil {
		fmt.Printf("\n❌ 请求失败: %v\n\n", err)
		return
	}

	fmt.Printf("\n🤖 [%s](%s):\n", model, currentProvider.Name())
	for chunk := range stream {
		fmt.Print(chunk)
	}
	fmt.Println()
}

// TODO: mermory
