package main

import (
	"fmt"
	"time"
)

func fu(c <-chan int) {
	time.Sleep(time.Second * 2)
	fmt.Println("after 2 second fu!")
	<-c
}
func main() {
	var done = make(chan int, 2)
	done <- 1
	done <- 2
	fmt.Println("", <-done)
	fmt.Println("", <-done)
	close(done)
	fmt.Println("len(done)", len(done))
	fmt.Println("cap(done)", cap(done))

}
