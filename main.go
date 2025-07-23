package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-example",
	Short: "一个极简的Cobra CLI应用程序示例",
	Long: `这是一个使用Cobra库创建的极简CLI应用程序示例。
它展示了Cobra的基本功能，包括命令、子命令和标志。`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示应用程序版本",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cobra-example v1.0.0")
	},
}

var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "向指定的人打招呼",
	Long:  `向指定的人发送友好的问候消息。如果没有提供名字，将使用默认问候。`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := "世界"
		if len(args) > 0 {
			name = args[0]
		}

		formal, _ := cmd.Flags().GetBool("formal")
		if formal {
			fmt.Printf("您好，%s！很高兴见到您。\n", name)
		} else {
			fmt.Printf("你好，%s！\n", name)
		}
	},
}

var echoCmd = &cobra.Command{
	Use:   "echo [message]",
	Short: "回显消息",
	Long:  `回显提供的消息，支持大写和重复选项。`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]

		uppercase, _ := cmd.Flags().GetBool("uppercase")
		repeat, _ := cmd.Flags().GetInt("repeat")

		if uppercase {
			message = strings.ToUpper(message)
		}

		for i := 0; i < repeat; i++ {
			fmt.Println(message)
		}
	},
}

func init() {
	// 添加子命令
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(echoCmd)

	// 为greet命令添加标志
	greetCmd.Flags().BoolP("formal", "f", false, "使用正式问候语")

	// 为echo命令添加标志
	echoCmd.Flags().BoolP("uppercase", "u", false, "将消息转换为大写")
	echoCmd.Flags().IntP("repeat", "r", 1, "重复消息的次数")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
