package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	ch := make(chan string, 10)
	go func() {
		for x := 0; x <= 100; x++ {
			if x%10 == 0 {
				fmt.Println("Generate UUID:")
				ch <- GenUUID()
				// time.Sleep(time.Millisecond)
			}

		}
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("async get:%v \n", v)
	}
}
func GenUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}
