package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "单次对话",
	Long:  `与大模型进行一次对话。`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("请输入问题")
			os.Exit(1)
		}
		initProvider()

		question := args[0]
		handleChat(question)
	},
}

func init() {
	chatCmd.Flags().StringVarP(&model, "model", "m", "", "模型名称")
	chatCmd.Flags().BoolVarP(&stream, "stream", "s", true, "启用流式输出")
	chatCmd.Flags().StringVarP(&provider, "provider", "p", "", "指定供应商 (volcengine, deepseek)")

	interactiveCmd.AddCommand(chatCmd)
}

func handleChat(question string) {
	// 调用API
	fmt.Printf("👤 用户: %s\n", question)
	fmt.Printf("🤖 [%s](%s): \n", model, currentProvider.Name())

	// 使用流式响应
	stream, err := currentProvider.CreateChatCompletion(question, model, true)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}

	// 打印流式响应
	for chunk := range stream {
		fmt.Print(chunk)
	}
	fmt.Println()
}
