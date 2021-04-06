package testcase02

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() error {
	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err:", err)
		return err
	}
	//保存到文件
	filePath := "D:\\test\\monster.json"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err:", err)
		return err
	}
	return err
}

func (this *Monster) ReStore() error {
	//从文件中读取序列化后的字符串
	filePath := "D:\\test\\monster.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err:", err)
		return err
	}

	//使用读取到的data []byte,进行反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return err
	}
	return err
}
