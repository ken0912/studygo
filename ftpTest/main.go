package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// fullpath := path.Join("D:\\Tools", "UploadFTP.done")
	// ok := Exists(fullpath)
	// if ok {
	// 	os.Remove(fullpath)
	// }

	// c, err := ftp.Dial("ftp2.salary.com:22", ftp.DialWithTimeout(5*time.Second))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = c.Login("sdcdba", "7jSdEWRH")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data := bytes.NewBufferString("Hello World")
	// err = c.Stor("/push_root/ConsumerSQL_new/JobBoard/20200714/test-file.txt", data)
	// if err != nil {
	// 	panic(err)
	// }
	// if err := c.Quit(); err != nil {
	// 	log.Fatal(err)
	// }
	var (
		err        error
		sftpClient *sftp.Client
	)

	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err = connect("sdcdba", "7jSdEWRH", "ftp2.salary.com", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	// 用来测试的本地文件路径 和 远程机器上的文件夹
	var localFilePath = "D:/Tools/UploadFTP.done"
	var remoteDir = "/push_root/ConsumerSQL_new/JobBoard/"
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	fmt.Println("copy file to remote server finished!")
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 5 * time.Second,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}
