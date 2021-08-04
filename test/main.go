package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var rowChan = make(chan int, 10)
	rowChan <- 1
	rowChan <- 2
	rowChan <- 3

	for i := range rowChan {
		fmt.Println("i:", i)
	}

	//RemoveDepulicates
	array := []int{0, 1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 5, 6, 6, 7, 7, 8}

	n := RemoveDepulicates(array)
	fmt.Println("n:", n)

	//StockMaxProfit
	prices := []int{7, 6, 4, 3, 1}
	MaxProfit := StockMaxProfit(prices)
	fmt.Println("MaxProfit:", MaxProfit)

	//rotate
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)

	fmt.Println("rand.Int63():", rand.Int63())
	fmt.Println("int64(len(charset)):", int64(8))

	var a uint64 = 11
	var b uint64 = 10
	fmt.Println(b - a)

	fmt.Println("--------------------------------------------------------")
	cnt := make(map[int]int, 2)
	for i := range random(100) {
		cnt[i] += 1
	}
	fmt.Println("cnt:", cnt)
}

func RemoveDepulicates(nums []int) int {
	low := 0
	length_nums := len(nums)
	for high := 1; high < length_nums; high++ {
		if nums[low] != nums[high] {
			nums[low+1] = nums[high]
			low++
		}
	}
	fmt.Println(nums[0 : low+1])
	return low + 1
}

func StockMaxProfit(prices []int) int {
	sumProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			sumProfit += prices[i+1] - prices[i]
		}
	}
	return sumProfit
}

func rotate(nums []int, k int) {
	/*
		arr1 := nums[len(nums)-k:]
		arr2 := nums[:len(nums)-k]
		nums = append(arr1, arr2...)
		fmt.Println("nums:", nums)
	*/
	reverse(nums, 0, len(nums)-1)
	fmt.Println("nums:", nums)
	reverse(nums, 0, k-1)
	fmt.Println("nums:", nums)
}
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			}
		}
	}()
	return c
}
