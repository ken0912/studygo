package main

import "fmt"

func BinarySearch(arr []int, num int) int {
	start := 0
	end := len(arr) - 1
	mid := (start + end) / 2
	for i := 0; i < len(arr); i++ {
		if num == arr[mid] {
			return mid
		}
		if num > arr[mid] {
			start = mid
		}
		if num < arr[mid] {
			end = mid
		}
		mid = (start + end) / 2
	}
	return -1
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := BinarySearch(arr, 9)
	fmt.Println(index)
}
