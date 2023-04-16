package deploy

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

type deployFlagInfo struct {
	host string
	src  string
	dest string
	key  string
}

var deployFlag deployFlagInfo

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "发布应用程序",
	Long:  `发布应用程序.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Subcommand!")
		if len(args) > 0 {
			fmt.Println("Args:", args)
		}

		copyFiles(cmd)
	},
}

func init() {
	DeployCmd.Flags().StringVarP(&deployFlag.host, "host", "h", "", "主机和端口 xxxxx.com:8888")
	DeployCmd.Flags().StringVarP(&deployFlag.key, "key", "k", "", "读取SSH私钥文件")
	DeployCmd.Flags().StringVarP(&deployFlag.src, "src", "s", "/zhangying", "来源目录")
	DeployCmd.Flags().StringVarP(&deployFlag.dest, "dest", "d", "/zhangying", "目标目录")
}

func copyFiles(cmd *cobra.Command) {
	host := cmd.Flags().Lookup("host").Value.String()
	src := cmd.Flags().Lookup("src").Value.String()
	dest := cmd.Flags().Lookup("dest").Value.String()

	// 上传目录
	sourceDir := src       // "/path/to/source/directory"
	destinationDir := dest // "/path/to/destination/directory"
	// 读取SSH私钥文件
	// privateKeyBytes, err := ioutil.ReadFile("/path/to/private/key")
	privateKeyBytes, err := ioutil.ReadFile(deployFlag.key)
	if err != nil {
		log.Fatal(err)
	}

	// 解析SSH私钥
	privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	// SSH客户端配置
	sshConfig := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接到Linux操作系统
	// sshClient, err := ssh.Dial("tcp", "linux-server:22", sshConfig)
	sshClient, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer sshClient.Close()

	// 创建SFTP客户端
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果是目录，则创建对应的目录
		if info.IsDir() {
			relativePath, _ := filepath.Rel(sourceDir, path)

			// 如果目标目录不存在，则创建目标目录
			_, err = sftpClient.Stat(filepath.Join(destinationDir, relativePath))
			if os.IsNotExist(err) {
				return sftpClient.MkdirAll(filepath.Join(destinationDir, relativePath))
			}
			if err != nil {
				return err
			}

			return nil
		}

		// 如果是文件，则上传文件
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		relativePath, _ := filepath.Rel(sourceDir, path)

		dstFile, err := sftpClient.Create(filepath.Join(destinationDir, relativePath))
		if os.IsExist(err) {
			dstFile, err = sftpClient.OpenFile(filepath.Join(destinationDir, relativePath), os.O_WRONLY|os.O_TRUNC)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}

		fmt.Println(relativePath, "上传成功！")

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("目录上传成功！")
}
