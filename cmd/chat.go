package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"ag/api"
	"ag/config"
)

var (
	model   string
	stream  bool
	apiKey  string
	rootCmd = &cobra.Command{ // 定义 rootCmd 变量
		Use:   "chat",
		Short: "AI 命令行工具",
		Long:  `ag 是一个与大模型交互的命令行工具。`,
		Run: func(cmd *cobra.Command, args []string) {
		if err := config.Load(); err != nil {
			fmt.Printf("加载配置失败: %v\n", err)
			os.Exit(1)
		}
		initProvider()
		startREPL()
	},
		
	}
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
	chatCmd.Flags().StringVarP(&model, "model", "m", "gpt-3.5", "模型名称")
	chatCmd.Flags().BoolVarP(&stream, "stream", "s", true, "启用流式输出")

	rootCmd.AddCommand(chatCmd)
}

func handleChat(question string) {
    // 获取默认提供商
    providerName := config.GetDefaultProvider()
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
