package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

//结构体序列化
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "1510-01-01",
		Sal:      50000,
		Skill:    "牛魔拳",
	}

	//将monster序列化json
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败:", err)
	}

	fmt.Println("monster序列化后为:", string(data))
}

//map序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火焰洞"

	//序列化json
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败:", err)
	}
	fmt.Println("a map 序列化的结果为:", string(data))
}

//切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = [5]string{"上海", "北京"}

	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 7
	m2["address"] = [5]string{"墨西哥", "洛杉矶"}

	slice = append(slice, m2)

	//序列化

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败:", err)
	}
	fmt.Println("slice序列化后结果为:", string(data))

}

//对基本数据类型序列化
func testFloat64() {
	var num1 float64 = 2345.67
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化失败:", err)
	}
	fmt.Println("num1序列化后结果为:", string(data))
}

func testunmarshal() {
	str := `{"name":"牛魔王","age":500,"birthday":"1510-01-01","sal":50000,"skill":"牛魔拳"}`
	var monster1 Monster
	err := json.Unmarshal([]byte(str), &monster1)
	if err != nil {
		fmt.Println("反序列化失败!", err)
	}
	fmt.Println("monster1反序列化结果为:", monster1)
}

func testunmarshalmap() {
	str := `{"name":"牛魔王","age":500,"birthday":"1510-01-01","sal":50000,"skill":"牛魔拳"}`
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化失败!", err)
	}
	fmt.Println("a反序列化结果为:", a)
}
func testunmarshalslice() {
	str := `[{"address":["上海","北京","","",""],"age":7,"name":"jack"},{"address":["墨西哥","洛杉矶","","",""],"age":7,"name":"tom"}]`
	var a []map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化失败!", err)
	}
	fmt.Println("a反序列化结果为:", a)
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
	testunmarshal()
	testunmarshalmap()
	testunmarshalslice()
}
