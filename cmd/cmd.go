// 命令行
// date: 2023-04-16
//
//auth:zhangying
package main

import (
	"fmt"
	"maya/cmd/deploy"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "maya",
	Short: "maya commond",
	Long:  `maya commond.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello,maya")
		// fmt.Println(args)
	},
}

type rootFlagInfo struct {
	userName string
	password string
}

var rootFlag rootFlagInfo

func init() {
	rootCmd.AddCommand(deploy.DeployCmd)

	rootCmd.PersistentFlags().StringVarP(&rootFlag.userName, "user", "u", "", "输入用户名")
	rootCmd.PersistentFlags().StringVarP(&rootFlag.password, "password", "p", "", "输入密码")

	// rootCmd.Flags().StringVarP(&r1, "root1", "r", "rootstring", "root flag string")
	// rootCmd.Flags().IntVarP(&r2, "root2", "t", 100, "root flag int")
	// subCmd.Flags().StringVarP(&s1, "sub1", "s", "substring", "A sub flag")
	// subCmd.Flags().IntVarP(&s2, "sub2", "i", 200, "A sub int flag")

	// rootCmd.Flags().Float32Var()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// TODO cmd 增加了linux远程文件拷贝功能，取代winscp操作，命令行一键操作，实际使用仍要完善修改
//  编码参考： https://cobra.dev/
