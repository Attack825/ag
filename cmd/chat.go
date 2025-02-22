package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	model   string
	stream  bool
	apiKey  string
	rootCmd = &cobra.Command{ // 定义 rootCmd 变量
		Use:   "ag",
		Short: "AI 命令行工具",
		Long:  `ag 是一个与大模型交互的命令行工具。`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "与大模型对话",
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
	// 在这里实现聊天命令的逻辑
	fmt.Printf("你问: %s\n", question)
	fmt.Println("模型回答: 这是模型的回答内容")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
