package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"ag/config"
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
    // 获取默认提供商
    cfg := config.GetProviderConfig(provider)
    if cfg == nil {
        fmt.Printf("找不到提供商配置: %s\n", provider)
        return
    }

	// 获取模型
    if model == "" {
		model = cfg.Model
    }

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
    fmt.Println() // 换行
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
