package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	cobra.OnInitialize(initConfig)
}

// 初始化命令行输出
var rootCmd = &cobra.Command{
	Use:   "plato",
	Short: "The reliable, scalable IM infrastructure for real-time communication.",
	Run:   Plato,
}

// Execute 执行函数
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Plato(cmd *cobra.Command, args []string) {

}

func initConfig() {

}
