package main

/*
	演示channel的同步作用
*/

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	quit := make(chan bool)
	go send(done, quit)
	go recv(done)

	//waiting for the task
	<-quit
	fmt.Println("main is over!")
}

func send(done, quit chan bool) {
	<-done
	fmt.Println("sent!")

	quit <- true
}

func recv(done chan bool) {
	time.Sleep(10 * time.Second)
	//Notify that the task is complete

	fmt.Println("func send can close")

	done <- true
}
