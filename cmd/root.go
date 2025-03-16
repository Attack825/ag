package cmd

import (
	"ag/api"
	"ag/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	model           string
	stream          bool
	provider        string
	currentProvider api.Provider
	rootCmd         = &cobra.Command{ // 定义 rootCmd 变量
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
	rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(chatCmd)
}

func initProvider() {
	// 从配置加载默认提供商
	var providerName string
	if provider != "" {
		providerName = provider // 使用命令行指定的供应商
	} else {
		providerName = config.GetDefaultProvider() // 使用默认供应商
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

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
