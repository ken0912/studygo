package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Num struct {
	Number int64
}
type Result struct {
	workid    int
	Num       *Num
	Sum       int64
	StartTime time.Time
	EndTime   time.Time
}

var wg sync.WaitGroup

func putNumRobot(numChan chan<- *Num) {
	var n = &Num{}

	for {
		n.Number = rand.Int63()
		// time.Sleep(time.Second * 2)
		numChan <- n
	}

}

func processRobot(numChan <-chan *Num, resultChan chan<- *Result, workid int) {
	for {
		var result = &Result{}
		var sum int64
		num := <-numChan
		start := time.Now()
		n := num.Number
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		r := rand.Intn(3)
		time.Sleep(time.Second * time.Duration(r))
		end := time.Now()
		result.Num = num
		result.Sum = sum
		result.workid = workid
		result.StartTime = start
		result.EndTime = end
		resultChan <- result
	}
}

func main() {
	numChan := make(chan *Num, 10)
	resultChan := make(chan *Result, 10)
	wg.Add(1)
	go putNumRobot(numChan)
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go processRobot(numChan, resultChan, i)
	}
	for r := range resultChan {
		fmt.Println(r.workid, r.Num.Number, r.Sum, r.StartTime, r.EndTime)
	}
	wg.Wait()

}
