package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("goroutine exit")
				return
			default:
				fmt.Println("goroutine running")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	stop <- true
	time.Sleep(2 * time.Second)
	fmt.Println("main fun exit")
}
