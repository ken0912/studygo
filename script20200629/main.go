package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	sourthpath := `\\salarysql26.salarynet.local\G$\Cleanup\2020-July\600_cleanup_physician\7_Cleanup_ByWeili`
	// sourthpath := `D:\Test\CompData\20200629`
	xfiles, _ := GetAllFiles(sourthpath)

	for _, v := range xfiles {

		content, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Println("err:", err)
			panic(err)
		}
		newcontent := string(content)
		var isreplace bool
		if strings.Contains(newcontent, "133") {
			newcontent = strings.Replace(newcontent, "133", "136", -1)
			isreplace = true
		}
		if strings.Contains(newcontent, "132") {
			newcontent = strings.Replace(newcontent, "132", "135", -1)
			isreplace = true
		}
		if isreplace {
			func(file string, content []byte) {
				nf, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Println("err:", file, err)
					panic(err)
				}
				defer nf.Close()
				_, err = nf.WriteString(newcontent)
				if err != nil {
					fmt.Println("err:", err)
					panic(err)
				}
			}(v, content)
			fmt.Printf("%v done.\n", v)
		}

	}
}

func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".sql")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}
