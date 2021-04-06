package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	// range 一直读取直到chan关闭，否则产生阻塞死锁
	//解决方式：
	//
	//a. 显式关闭 channel；
	//
	//b. for range chan放进子协程，主协程 sleep 等待时间后退出；
	func() {
		for {
			v, ok := <-ch
			if !ok {
				return
			} else {
				fmt.Println("v:", v)
			}
		}
	}()

}
