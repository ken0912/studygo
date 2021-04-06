package main

import (
	"fmt"
	"os"
)

func main() {
	filepath := `abc.txt`
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file error!", err)
	}
	defer file.Close()

	fmt.Println("file:", file)
}
