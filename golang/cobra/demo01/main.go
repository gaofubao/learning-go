package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var echoTimes int

	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1), // 3. Args（或Flag）校验
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		}, // 4. 编写命令响应代码
	}

	var cmdTimes = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "echo anything to the screen more times",
		Long:  `Echo things multiple times back to the user by providing a count and a string.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}

	// 2. 为命令指定参数
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	// 1. 定义命令对象
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdEcho)
	cmdEcho.AddCommand(cmdTimes)

	// 5. 执行命令
	rootCmd.Execute()
}
