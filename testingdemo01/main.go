package main

import "fmt"

func addUpper(n int) int {
	res := 0
	for i := 1; i < n; i++ {
		res += i
	}
	return res
}

func main() {
	res := addUpper(10)
	fmt.Println("res:", res)
	if res == 55 {
		fmt.Printf("addUpper期望值%d,返回值%d,正确!", 55, res)
	} else {
		fmt.Printf("addUpper期望值%d,返回值%d,错误!", 55, res)
	}
}
