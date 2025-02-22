package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"ag/api"
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

		question := args[0]
		handleChat(question)
	},
}

func init() {

	interactiveCmd.AddCommand(chatCmd)
}

func handleChat(question string) {
    // 获取默认提供商
    var providerName string
	if provider != "" {
		providerName = provider  // 使用命令行指定的供应商
	} else {
		providerName = config.GetDefaultProvider()  // 使用默认供应商
	}
	
	if providerName == "" {
		fmt.Println("未配置默认提供商")
		return
	}

    // 获取提供商实例
    provider := api.GetProvider(providerName)
    if provider == nil {
        fmt.Printf("找不到提供商: %s\n", providerName)
        return
    }

    // 调用API
    fmt.Printf("用户: %s\n", question)
    fmt.Printf("%s回答: \n", provider.Name())
    
    // 使用流式响应
    stream, err := provider.CreateChatCompletion(question, true)
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
