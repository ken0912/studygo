package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := sqrt(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", result, err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math:square root of negative number")
	}

	return 0, nil
}
