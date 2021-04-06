package sort

//The Developer only needs to implement the following interface to use the following sort method.
type Vector interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//BubbleSort
func BubbleSort(vector Vector) Vector {
	l := vector.Len()
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if vector.Less(j, j-1) {
				vector.Swap(j, j-1)
			}
		}
	}
	return vector
}

//SelectSort
func SelectSort(vector Vector) Vector {
	l := vector.Len()
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if vector.Less(j, min) {
				min = j
			}
		}
		vector.Swap(i, min)
	}
	return vector
}
