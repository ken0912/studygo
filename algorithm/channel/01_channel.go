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
				fmt.Println("monitor quite")
				return
			default:
				fmt.Println("goroutine monitoring")
				time.Sleep(1 * time.Second)
			}

		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true

	//To detect if the monitoring has stopped
	time.Sleep(5 * time.Second)

}
