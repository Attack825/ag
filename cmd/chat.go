package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	model  string
	stream bool
	apiKey string
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
