package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		datapath := `D:\Study\GO\src\github.com\ken0912\studygo\test20200622\test.exe`
		cmd := exec.Command("cmd.exe", "/c", "start "+datapath)
		fmt.Println("starts")
		cmd.Run()
		fmt.Println("end")
		currentpath, _ := os.Getwd()
		fmt.Println("Getwd:", currentpath)
	*/
	/*
		layout := "2006-01-02"
		strdate := "2020-06-38"

		cdate, err := time.Parse(layout, strdate)
		fmt.Println("cdate:", cdate)
		fmt.Println("err:", err)
	*/
	data := make(chan int, 3)
	exit := make(chan bool)
	var a int
	go func() {
		// for d := range data {
		// 	fmt.Println(d)
		// }
		for {
			a = <-data
			fmt.Println(a)
		}
		exit <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			data <- i
		}
		close(data)
	}()

	<-exit
}
