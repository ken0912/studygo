package main

import "fmt"

func BubbleSort_1(arr []int) {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	fmt.Println("count:", count)
}

//冒泡排序优化
func BubbleSort2(arr []int) {
	count := 0
	flg := false
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flg = true
			}
		}
		fmt.Println("flg:", flg)
		if !flg {
			fmt.Println("count", count)
			return
		} else {
			flg = false
		}
	}
}

func main() {
	arr := []int{1, 3, 8, 12, 2, 6, 7, 8, 5, 4, 3}
	// arr := []int{1, 2, 3, 4, 5, 6, 7}
	// BubbleSort_1(arr)
	BubbleSort2(arr)

	fmt.Println("arr:", arr)

}
