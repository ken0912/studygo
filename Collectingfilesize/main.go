package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var isProcessSubDir = true
var TotalSize int64

func main() {
	_, _ = GetAllFiles(`C:\`)
	fmt.Println("TotalSize:", float64(TotalSize)/1024/1024, "MB")
}

func GetAllFiles(dirPth string) (files []string, err error) {
	// var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() && isProcessSubDir { // 目录, 递归遍历
			// dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// // 过滤指定格式
			// ok := strings.HasSuffix(fi.Name(), fileType)
			// if ok {
			TotalSize += fi.Size() /// 1024.0 / 1024.0
			// fmt.Println(dirPth+PthSep+fi.Name(), ":", TotalSize)
			// files = append(files, dirPth+PthSep+fi.Name())
			// }
		}
	}

	// 读取子目录下文件
	// if isProcessSubDir {
	// 	for _, table := range dirs {
	// 		temp, _ := GetAllFiles(table)
	// 		for _, temp1 := range temp {
	// 			// files = append(files, temp1)
	// 		}
	// 	}
	// }

	return files, nil
}
