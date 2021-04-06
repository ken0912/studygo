package main

import "fmt"

func main() {
	ret := step(10)
	fmt.Println("ret:", ret)
}

/*

 */
func step(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return step(n-1) + step(n-2)
}
