package sort

import (
	"fmt"
	"testing"
)

type List []int

func (l List) Len() int {
	return len(l)
}
func (l List) Less(i, j int) bool {
	return l[i] > l[j]
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func Test_Xxxxx(t *testing.T) {
	l := List{3, 9, 2, 6, 5, 1, 7, 34, 3, 54, 53, 45, 345, 4, 34, 745, 7, 65, 85, 8, 67, 9867, 9, 67, 9}
	fmt.Println(BubbleSort(l))
}

func Test_Selectsort(t *testing.T) {
	l := List{3, 9, 2, 6, 5, 1, 7, 34, 3, 54, 53, 45, 345, 4, 34, 745, 7, 65, 85, 8, 67, 9867, 9, 67, 9}
	fmt.Println(SelectSort(l))
}

func Test_XXxxx(t *testing.T) {
	fmt.Println("begin")
	for i := 1; i <= 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println("i=", i)
	}

	fmt.Println("end")
}
