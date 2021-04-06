package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())
		// \\10.100.10.51\\g$\\EnterpriseSQL\\IPAS\\GSYBcpDataStoreForDailyUpdate
		fullpath := path.Join(`\\10.100.10.51\g$\EnterpriseSQL\IPAS\GSYBcpDataStoreForDailyUpdate`, "UploadToFTP.done")
		ok := Exists(fullpath)
		if ok {
			fmt.Println("exists!")
		}
		fmt.Println("fullpath:", fullpath)
	}
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
