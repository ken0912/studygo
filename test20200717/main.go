package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.IsExist(nil))
	s := generateString(0, 30)
	fmt.Println("s:", s)
}

func Exists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func generateString(x int, n int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		fmt.Println("a:", i%len(letters))
		b[i] = letters[i]
	}
	return string(b)
}
