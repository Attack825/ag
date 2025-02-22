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

var currentProvider api.Provider

func init() {
	rootCmd.AddCommand(interactiveCmd)
}

var interactiveCmd = &cobra.Command{
	Use:   "ag",
	Short: "进入交互式对话模式",
	
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
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("[%s] 输入问题（输入 exit 退出）:\n", currentProvider.Name())

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
	stream, _ := currentProvider.CreateChatCompletion(prompt, true)
	fmt.Printf("\n[%s] 回答:\n", currentProvider.Name())
	for chunk := range stream {
		fmt.Print(chunk)
	}
	fmt.Println("\n")
}
