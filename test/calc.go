package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

func main() {
	// s1 := []int{1, 2, 3}
	// s2 := s1[1:2]
	// s3 := append(s2, 10, 11)
	// fmt.Println("s1:", s1, &s1)
	// fmt.Println("s2:", s2, &s2)
	// fmt.Println("s3:", s3, &s3)
	s := 1 << 20
	fmt.Println("s:", s)
}
