package main

import (
	"fmt"
	"math/rand"
)

func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		index := 0
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] > arr[index] {
				index = j
			}
		}
		//每次找到最大的值 放到末尾
		arr[len(arr)-1-i], arr[index] = arr[index], arr[len(arr)-1-i]
	}
}
func main() {
	arr := []int{12, 2, 5, 6, 7, 9, 3, 4}
	SelectSort(arr)
	fmt.Println("arr:", arr)

	numList := rand.Perm(10)
	fmt.Println("numList:", numList)
}
