package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	n := 8
	str := make([]byte, n)
	fmt.Println("rand.Int63():", rand.Int63())
	fmt.Println("int64(len(charset)):", int64(len(charset)))
	for i, _ := range str {
		str[i] = charset[rand.Int63()%int64(len(charset))]
	}
	fmt.Println("str:", string(str))
}
