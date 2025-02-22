package cmd

import (
	"ag/api"
	"ag/config"
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	model   string
	stream  bool
	provider string
	currentProvider api.Provider
	rootCmd = &cobra.Command{ // 定义 rootCmd 变量
		Use:   "ag",
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

func init() {
	interactiveCmd.Flags().StringVarP(&model, "model", "m", "", "模型名称")
	interactiveCmd.Flags().BoolVarP(&stream, "stream", "s", true, "启用流式输出")
	interactiveCmd.Flags().StringVarP(&provider, "provider", "p", "", "指定供应商 (volcengine, deepseek)")  // 新增 provider 参数

	rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(chatCmd)
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "进入交互式对话模式",
	Run: func(cmd *cobra.Command, args []string) {
		initProvider()
		startREPL()
	},
}

func initProvider() {
    // 从配置加载默认提供商
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
    
    currentProvider = api.GetProvider(providerName)
    if currentProvider == nil {
        fmt.Printf("找不到提供商: %s\n", providerName)
        os.Exit(1)
    }

	// 获取模型
    if model == "" {
        // 如果命令行未指定模型，使用配置中的默认模型
        if cfg := config.GetProviderConfig(providerName); cfg != nil {
            model = cfg.Model
        }
    }

	// currentProvider.SetModel(modelName)
}

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
    fmt.Println("\n")
}
