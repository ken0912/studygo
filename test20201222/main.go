package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go producer(ch)
	go producer1(ch)
	for v := range ch {
		fmt.Println("v:", v)
	}
}
func producer(ch chan string) {
	str := "abcdefg"
	for _, v := range str {
		ch <- string(v)
	}

}
func producer1(ch chan string) {
	str := "hijklmn"
	for _, v := range str {
		ch <- string(v)
	}
}
